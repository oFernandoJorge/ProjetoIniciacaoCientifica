import { useMemo, useState, type ChangeEvent } from "react";
import * as XLSX from "xlsx";
import { api } from "../services/api";

type SubmissionRow = {
  title: string;
  presenterName: string;
  course: string;
  knowledgeArea: string;
  modality: string;
  campus: string;
  advisorName: string;
  presentationType: string;
};

type Room = {
  name: string;
  floor: number;
  presentationType: string;
  capacity: number;
};

type SubmissionGroup = {
  knowledgeArea: string;
  presentationType: string;
  courses: string[];
  submissions: SubmissionRow[];
};

type RoomAllocation = {
  roomName: string;
  floor: number;
  knowledgeArea: string;
  presentationType: string;
  submissions: SubmissionRow[];
};

function normalizeText(value: string) {
  return String(value || "")
    .toUpperCase()
    .replace(/[.]/g, "") // Remove pontos
    .trim();
}

function parseSpreadsheetRows(rows: unknown[][]): SubmissionRow[] {
  const submissions: SubmissionRow[] = [];

  for (let index = 0; index < rows.length; index += 1) {
    const row = rows[index];

    if (index === 0) {
      continue;
    }

    if (!row || row.length < 25) {
      continue;
    }

    // Mapeamento baseado na análise da planilha real:
    // 1: Título
    // 5: Nome do apresentador
    // 23: Cursos de vinculação/formação
    // 22: Área de conhecimento
    // 10: Modalidade de apresentação
    // 24: Local de apresentação (Campus)
    // 2: Autor(es) (usado como advisorName por falta de coluna específica)
    
    submissions.push({
      title: String(row[1] ?? "").trim(),
      presenterName: String(row[5] ?? "").trim(),
      course: normalizeText(String(row[23] ?? "")),
      knowledgeArea: normalizeText(String(row[22] ?? "")),
      modality: String(row[10] ?? "").trim(),
      campus: String(row[24] ?? "").trim(),
      advisorName: String(row[2] ?? "").trim(),
      presentationType: String(row[10] ?? "").trim(),
    });
  }

  return submissions;
}

function generateRooms(): Room[] {
  const rooms: Room[] = [];

  for (let floor = 2; floor <= 5; floor += 1) {
    // Salas 201 a 210: E-POSTER (5 min)
    for (let roomNumber = 1; roomNumber <= 10; roomNumber += 1) {
      rooms.push({
        name: `Sala ${floor}${roomNumber.toString().padStart(2, "0")}`,
        floor,
        presentationType: "E-POSTER",
        capacity: 12, // Ex: 12 apresentações de 5 min em 1 hora
      });
    }

    // Salas 211 a 220: ORAL (20 min)
    for (let roomNumber = 11; roomNumber <= 20; roomNumber += 1) {
      rooms.push({
        name: `Sala ${floor}${roomNumber.toString().padStart(2, "0")}`,
        floor,
        presentationType: "ORAL",
        capacity: 6, // Ex: 6 apresentações de 20 min em 2 horas
      });
    }
  }

  return rooms;
}

function normalizePresentationType(value: string) {
  const normalized = String(value || "")
    .toUpperCase()
    .normalize("NFD")
    .replace(/[\u0300-\u036f]/g, "") // Remove acentos (ex: PÔSTER -> POSTER)
    .replace(/\s+/g, " ")
    .trim();

  if (normalized.includes("POSTER")) {
    return "E-POSTER";
  }

  if (normalized.includes("ORAL")) {
    return "ORAL";
  }

  return "ORAL"; // Default seguro
}

function formatScheduleTime(index: number, turn: string) {
  let startHour = 8;
  const normalizedTurn = String(turn || "").toUpperCase();
  
  if (normalizedTurn.includes("TARDE")) {
    startHour = 14;
  } else if (normalizedTurn.includes("NOITE")) {
    startHour = 19;
  }

  const baseMinutes = 0 + index * 20;
  const hour = startHour + Math.floor(baseMinutes / 60);
  const minutes = baseMinutes % 60;
  return `${hour.toString().padStart(2, "0")}:${minutes.toString().padStart(2, "0")}`;
}

function contains(values: string[], target: string) {
  return values.some((value) => value === target);
}

function groupSubmissionsByArea(submissions: SubmissionRow[]): SubmissionGroup[] {
  const groupsMap: Record<string, SubmissionGroup> = {};

  for (const submission of submissions) {
    const presentationType = normalizePresentationType(submission.presentationType);
    const knowledgeArea = submission.knowledgeArea || "SEM ÁREA";
    const key = `${knowledgeArea}_${presentationType}`;

    if (!groupsMap[key]) {
      groupsMap[key] = {
        knowledgeArea,
        presentationType,
        courses: [],
        submissions: [],
      };
    }

    const group = groupsMap[key];

    if (submission.course && !contains(group.courses, submission.course)) {
      if (group.courses.length < 3) {
        group.courses.push(submission.course);
      }
    }

    group.submissions.push({
      ...submission,
      presentationType,
      knowledgeArea,
    });
  }

  return Object.values(groupsMap);
}

function allocateRooms(groups: SubmissionGroup[], rooms: Room[]): RoomAllocation[] {
  const allocations: RoomAllocation[] = [];
  let roomIndex = 0;

  for (const group of groups) {
    const compatibleRooms = rooms.filter(
      (room) => room.presentationType === group.presentationType,
    );

    if (compatibleRooms.length === 0) {
      continue;
    }

    let i = 0;
    while (i < group.submissions.length) {
      // Pega a próxima sala disponível
      const selectedRoom = compatibleRooms[roomIndex % compatibleRooms.length];
      const capacity = selectedRoom.capacity;
      
      const end = Math.min(i + capacity, group.submissions.length);

      allocations.push({
        roomName: selectedRoom.name,
        floor: selectedRoom.floor,
        knowledgeArea: group.knowledgeArea,
        presentationType: group.presentationType,
        submissions: group.submissions.slice(i, end),
      });

      // Pula para o próximo bloco de alunos
      i += capacity;
      // Pula para a PRÓXIMA sala
      roomIndex += 1;
    }
  }

  return allocations;
}

export function Upload() {
  const [file, setFile] = useState<File | null>(null);
  const [allocations, setAllocations] = useState<RoomAllocation[]>([]);
  const [warning, setWarning] = useState<string | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [message, setMessage] = useState<string | null>(null);
  const [scheduleDate, setScheduleDate] = useState("22/10/2025");
  const [scheduleTurn, setScheduleTurn] = useState("MANHÃ");
  const [scheduleCampus, setScheduleCampus] = useState("Campus Carneiro da Cunha");

  const rooms = useMemo(() => generateRooms(), []);

  const handleFileChange = (event: ChangeEvent<HTMLInputElement>) => {
    setError(null);
    setMessage(null);
    setWarning(null);
    setAllocations([]);
    setFile(event.target.files?.[0] ?? null);
  };

  const handleProcess = async () => {
    if (!file) {
      setError("Escolha um arquivo de planilha antes de processar.");
      return;
    }

    setError(null);
    setWarning(null);
    setMessage(null);

    try {
      const arrayBuffer = await file.arrayBuffer();
      const workbook = XLSX.read(arrayBuffer, { type: "array" });
      const sheetName = workbook.SheetNames[0];

      if (!sheetName) {
        throw new Error("A planilha não possui abas válidas.");
      }

      const rows = XLSX.utils.sheet_to_json<unknown[]>(workbook.Sheets[sheetName], {
        header: 1,
        raw: false,
      });

      const submissions = parseSpreadsheetRows(rows);

      if (submissions.length === 0) {
        throw new Error("Nenhuma linha de submissão válida encontrada na planilha.");
      }

      const groups = groupSubmissionsByArea(submissions);
      const allocationsResult = allocateRooms(groups, rooms);

      if (allocationsResult.length === 0) {
        setWarning("Nenhuma sala compatível foi encontrada para os tipos de apresentação informados.");
      }

      setAllocations(allocationsResult);
      setMessage(`Processadas ${submissions.length} submissões e alocadas ${allocationsResult.length} salas.`);
    } catch (err) {
      setAllocations([]);
      setError(err instanceof Error ? err.message : "Erro ao processar a planilha.");
    }
  };

  const handleDownloadPdf = async () => {
    if (allocations.length === 0) {
      setError("Não há alocações para gerar PDF.");
      return;
    }

    setError(null);

    const payload = {
      date: scheduleDate,
      turn: scheduleTurn,
      campus: scheduleCampus,
      items: allocations.map((allocation) => ({
        roomName: allocation.roomName,
        knowledgeArea: allocation.knowledgeArea,
        presentationType: allocation.presentationType,
        courses: allocation.submissions
          .map((submission) => submission.course)
          .filter((course, index, self) => course && self.indexOf(course) === index),
        submissions: allocation.submissions.map((submission, index) => ({
          time: formatScheduleTime(index, scheduleTurn),
          title: submission.title,
          presenterName: submission.presenterName,
        })),
      })),
    };

    try {
      const response = await api.post("/pdf", payload, {
        responseType: "blob",
      });

      const blob = new Blob([response.data], { type: "application/pdf" });
      const url = window.URL.createObjectURL(blob);

      const link = document.createElement("a");
      link.href = url;
      link.target = "_blank";
      link.download = "organizador-de-salas.pdf";
      document.body.appendChild(link);
      link.click();
      document.body.removeChild(link);

      setTimeout(() => window.URL.revokeObjectURL(url), 100);
    } catch (err) {
      setError("Erro ao gerar ou baixar o PDF. Verifique se o backend está rodando.");
      console.error(err);
    }
  };

  const maxStudents = useMemo(() => {
    return Math.max(...allocations.map(a => a.submissions.length), 1);
  }, [allocations]);

  return (
    <main className="app-shell">
      <div className="page-container">
        <header className="page-header">
          <p className="eyebrow">Sistema de Ensalamento</p>
          <h1>Organizador de Apresentações</h1>
          <p className="lead">
            Gerencie o ensalamento acadêmico de forma automatizada, respeitando áreas de conhecimento e turnos.
          </p>
        </header>

        <div className="card">
          <div className="card-header">
            <h3>Configurações e Upload</h3>
            <span className="badge">Excelize Engine</span>
          </div>

          <div className="form-grid">
            <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '20px' }}>
              <label>
                Selecione a planilha (.xlsx)
                <input type="file" accept=".xlsx,.xls,.csv" onChange={handleFileChange} />
              </label>

              <label>
                Data do Evento
                <input type="text" value={scheduleDate} onChange={(e) => setScheduleDate(e.target.value)} />
              </label>
            </div>

            <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '20px' }}>
              <label>
                Turno
                <input type="text" value={scheduleTurn} onChange={(e) => setScheduleTurn(e.target.value)} />
              </label>

              <label>
                Campus
                <input type="text" value={scheduleCampus} onChange={(e) => setScheduleCampus(e.target.value)} />
              </label>
            </div>

            <div style={{ marginTop: '12px' }}>
              <button type="button" onClick={handleProcess} className="primary-button" style={{ width: '100%' }}>
                Processar e Organizar Salas
              </button>
            </div>

            {message && <p className="message success">{message}</p>}
            {warning && <p className="message error">{warning}</p>}
            {error && <p className="message error">{error}</p>}
          </div>
        </div>

        {allocations.length > 0 && (
          <div className="card" style={{ textAlign: 'center' }}>
            <div className="card-header">
              <h3>Documentação Pronta</h3>
              <p className="lead">Tudo pronto para gerar o PDF oficial do evento.</p>
            </div>

            <button type="button" onClick={handleDownloadPdf} className="secondary-button" style={{ width: '100%', marginBottom: '32px' }}>
              Baixar Grade de Apresentações (PDF)
            </button>

            <div className="chart-container">
              <h3 style={{ marginBottom: '24px', fontSize: '1rem' }}>Distribuição de Alunos por Sala</h3>
              <div className="chart-bars">
                {allocations.map((allocation, i) => (
                  <div key={i} className="chart-bar-wrapper">
                    <div 
                      className="chart-bar" 
                      style={{ height: `${(allocation.submissions.length / maxStudents) * 100}%` }}
                      title={`${allocation.roomName}: ${allocation.submissions.length} alunos`}
                    >
                      <span className="chart-bar-value">{allocation.submissions.length}</span>
                    </div>
                    <span className="chart-bar-label">{allocation.roomName.replace("Sala ", "")}</span>
                  </div>
                ))}
              </div>
            </div>
          </div>
        )}
      </div>
    </main>
  );
}

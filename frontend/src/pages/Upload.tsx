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
      course: String(row[23] ?? "").trim(),
      knowledgeArea: String(row[22] ?? "").trim(),
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
    for (let roomNumber = 1; roomNumber <= 10; roomNumber += 1) {
      rooms.push({
        name: `Sala ${floor}0${roomNumber}`,
        floor,
        presentationType: "E-POSTER",
        capacity: 12,
      });
    }

    for (let roomNumber = 11; roomNumber <= 20; roomNumber += 1) {
      rooms.push({
        name: `Sala ${floor}${roomNumber}`,
        floor,
        presentationType: "ORAL",
        capacity: 6,
      });
    }
  }

  return rooms;
}

function normalizePresentationType(value: string) {
  const normalized = String(value || "")
    .toUpperCase()
    .replace(/\s+/g, " ")
    .replace(/[-_]/g, " ")
    .trim();

  if (normalized.includes("POSTER")) {
    return "E-POSTER";
  }

  if (normalized.includes("ORAL")) {
    return "ORAL";
  }

  return normalized || "UNKNOWN";
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

    const selectedRoom = compatibleRooms[roomIndex % compatibleRooms.length];
    const capacity = selectedRoom.capacity;

    for (let i = 0; i < group.submissions.length; i += capacity) {
      const end = Math.min(i + capacity, group.submissions.length);

      allocations.push({
        roomName: selectedRoom.name,
        floor: selectedRoom.floor,
        knowledgeArea: group.knowledgeArea,
        presentationType: group.presentationType,
        submissions: group.submissions.slice(i, end),
      });

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

      // Abre o PDF em uma nova aba
      const link = document.createElement("a");
      link.href = url;
      link.target = "_blank";
      // Opcional: define um nome para o arquivo se o navegador decidir baixar
      link.download = "organizador-de-salas.pdf";
      document.body.appendChild(link);
      link.click();
      document.body.removeChild(link);

      // Limpa a URL da memória após um tempo
      setTimeout(() => window.URL.revokeObjectURL(url), 100);
    } catch (err) {
      setError("Erro ao gerar ou baixar o PDF. Verifique se o backend está rodando.");
      console.error(err);
    }
  };

  return (
    <main className="page-container">
      <header className="page-header">
        <div>
          <p className="eyebrow">Organizador de Ensalamento</p>
          <h1>Envie a planilha do outro sistema</h1>
          <p className="lead">
            O sistema irá ler a planilha, agrupar por área de conhecimento e tipo de apresentação, e alocar cada grupo em salas.
          </p>
        </div>
      </header>

      <section className="page-grid">
        <article className="card">
          <div className="card-header">
            <h2>Controle de planilha</h2>
            <span className="badge">Formato esperado: 8 colunas</span>
          </div>

          <div className="form-grid">
            <label>
              Selecione a planilha
              <input type="file" accept=".xlsx,.xls,.csv" onChange={handleFileChange} />
            </label>

            <label>
              Data da sessão
              <input type="text" value={scheduleDate} onChange={(event) => setScheduleDate(event.target.value)} />
            </label>

            <label>
              Turno
              <input type="text" value={scheduleTurn} onChange={(event) => setScheduleTurn(event.target.value)} />
            </label>

            <label>
              Campus
              <input type="text" value={scheduleCampus} onChange={(event) => setScheduleCampus(event.target.value)} />
            </label>

            <div className="form-actions">
              <button type="button" onClick={handleProcess} className="primary-button">
                Organizar planilha
              </button>
            </div>

            {message && <p className="message success">{message}</p>}
            {warning && <p className="message error">{warning}</p>}
            {error && <p className="message error">{error}</p>}
          </div>
        </article>

        <article className="card">
          <div className="card-header">
            <h2>Resultado de alocação</h2>
            {allocations.length > 0 && <span className="badge">Salas geradas: {allocations.length}</span>}
          </div>

          {allocations.length === 0 ? (
            <div className="empty-state">
              Faça upload da planilha e clique em "Organizar planilha" para preparar o PDF.
            </div>
          ) : (
            <div className="actions-column">
              <p className="lead" style={{ textAlign: 'center', marginBottom: '1.5rem' }}>
                A alocação foi concluída com sucesso. Clique no botão abaixo para gerar o documento final.
              </p>
              <button 
                type="button" 
                onClick={handleDownloadPdf} 
                className="secondary-button"
                style={{ width: '100%', padding: '1.5rem', fontSize: '1.2rem' }}
              >
                Gerar PDF das Salas
              </button>
            </div>
          )}
        </article>
      </section>
    </main>
  );
}

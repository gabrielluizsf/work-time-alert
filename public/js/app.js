let timer;
let isRunning = false;
let currentSeconds = 0;
let defaultTime = 0;
let totalSeconds = 0;
let activeAlert = null;
let editingId = null;
let language = "pt";
let sessionId = null;

const SNOOZE_SECONDS = 5 * 60;
const RING_RADIUS = 124;
const RING_CIRCUMFERENCE = 2 * Math.PI * RING_RADIUS;

let audioCtx = null;
let alertSoundBuffer = null;
const display = document.getElementById("timer-display");
const ring = document.getElementById("ring-value");
const pageTitle = document.getElementById("page-title");
const timerCard = document.querySelector(".timer-card");
const alertsList = document.getElementById("alerts-list");
const alertsCount = document.getElementById("alerts-count");
const emptyState = document.getElementById("empty-state");
const currentTimeDisplay = document.getElementById("current-time");
const startBtn = document.getElementById("start-btn");
const pauseBtn = document.getElementById("pause-btn");
const resetBtn = document.getElementById("reset-btn");
const alertForm = document.getElementById("alert-form");
const nameInput = document.getElementById("alert-name");
const hoursInput = document.getElementById("alert-hours");
const minutesInput = document.getElementById("alert-minutes");
const secondsInput = document.getElementById("alert-seconds");
const cancelFormBtn = document.getElementById("cancel-form-btn");
const snackbar = document.getElementById("snackbar");
const dialog = document.getElementById("done-dialog");
const dialogTitle = document.getElementById("dialog-title");
const dialogMessage = document.getElementById("dialog-message");
const dialogOk = document.getElementById("dialog-ok");
const dialogSnooze = document.getElementById("dialog-snooze");
const addAlertLabel = document.getElementById("add-alert-label");
const cancelFormLabel = document.getElementById("cancel-form-label");

ring.style.strokeDasharray = RING_CIRCUMFERENCE;

const translations = {
    pt: {
        tagline: "Seu parceiro de foco no trabalho",
        activeLabel: "Alerta ativo",
        alertsTitle: "Meus Alertas",
        start: "Iniciar",
        pause: "Pausar",
        reset: "Resetar",
        addAlert: "Adicionar",
        cancel: "Limpar",
        cancelEdit: "Cancelar",
        nameLabel: "Nome do alerta",
        hoursLabel: "Horas",
        minutesLabel: "Min",
        secondsLabel: "Seg",
        use: "Usar",
        edit: "Editar",
        save: "Salvar",
        delete: "Excluir",
        undo: "Desfazer",
        ok: "OK",
        snooze: "Adiar 5 min",
        attention: "Atenção",
        invalidAlert: "Por favor, insira um nome válido e um tempo maior que 0.",
        timeUpMessage: "O tempo do seu alerta finalizou!",
        currentTime: "Horário atual",
        hint: "Dica: aperte Espaço para iniciar ou pausar o contador",
        added: "Alerta adicionado",
        updated: "Alerta atualizado",
        deleted: "Alerta excluído",
        snoozed: "Alerta adiado em 5 min",
        emptyTitle: "Nenhum alerta ainda",
        emptyDesc: "Crie seu primeiro alerta no formulário acima.",
    },
    en: {
        tagline: "Your work focus companion",
        activeLabel: "Active alert",
        alertsTitle: "My Alerts",
        start: "Start",
        pause: "Pause",
        reset: "Reset",
        addAlert: "Add",
        cancel: "Clear",
        cancelEdit: "Cancel",
        nameLabel: "Alert name",
        hoursLabel: "Hours",
        minutesLabel: "Min",
        secondsLabel: "Sec",
        use: "Use",
        edit: "Edit",
        save: "Save",
        delete: "Delete",
        undo: "Undo",
        ok: "OK",
        snooze: "Snooze 5 min",
        attention: "Attention",
        invalidAlert: "Please enter a valid name and a time greater than 0.",
        timeUpMessage: "Your alert's time has finished!",
        currentTime: "Current time",
        hint: "Tip: press Space to start or pause the countdown",
        added: "Alert added",
        updated: "Alert updated",
        deleted: "Alert deleted",
        snoozed: "Alert snoozed by 5 min",
        emptyTitle: "No alerts yet",
        emptyDesc: "Create your first alert using the form above.",
    },
};

const ICONS = {
    use: '<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"/><path d="M10 8v8l6-4z"/></svg>',
    edit: '<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M17 3a2.85 2.83 0 1 1 4 4L7.5 20.5 2 22l1.5-5.5L17 3Z"/></svg>',
    delete: '<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M3 6h18"/><path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6"/><path d="M8 6V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"/></svg>',
};

let alerts = [
    { id: 1, name: "Hora de trabalhar", hours: 0, minutes: 25, seconds: 0 },
    { id: 2, name: "Pausa Curta", hours: 0, minutes: 5, seconds: 0 },
];

function t(key) {
    return (translations[language] && translations[language][key]) || key;
}

function applyLanguage() {
    document.querySelectorAll("[data-i18n]").forEach((el) => {
        el.textContent = t(el.dataset.i18n);
    });
    document.querySelectorAll("[data-i18n-label]").forEach((el) => {
        el.textContent = t(el.dataset.i18nLabel);
    });
}

const sessions = [];
initSession(sessions).then(() => {
    if (sessions.length > 0) {
        sessionId = sessions[0];
        registerServiceWorker(sessionId);
    }
});

function formatTime(seconds) {
    const h = Math.floor(seconds / 3600).toString().padStart(2, "0");
    const m = Math.floor((seconds % 3600) / 60).toString().padStart(2, "0");
    const s = (seconds % 60).toString().padStart(2, "0");
    return `${h}:${m}:${s}`;
}

function totalOf(alert) {
    return alert.hours * 3600 + alert.minutes * 60 + alert.seconds;
}

function updateDisplay() {
    display.textContent = formatTime(currentSeconds);
    updateRing();
    updateControls();
}

function updateRing() {
    const fraction =
        totalSeconds > 0 ? clamp(currentSeconds / totalSeconds, 0, 1) : 1;
    ring.style.strokeDashoffset = RING_CIRCUMFERENCE * (1 - fraction);
    ring.style.stroke =
        fraction <= 0.1
            ? "var(--error)"
            : fraction <= 0.35
              ? "var(--warning)"
              : "var(--primary)";
}

function clamp(value, min, max) {
    return Math.min(Math.max(value, min), max);
}

function updateControls() {
    startBtn.disabled = isRunning || currentSeconds <= 0;
    pauseBtn.disabled = !isRunning;
    resetBtn.disabled = !isRunning && currentSeconds === defaultTime;
    timerCard.classList.toggle("is-running", isRunning);
}

function startTimer() {
    if (isRunning || currentSeconds <= 0) return;
    const ctx = getAudioContext();
    if (ctx) loadAlertSound(ctx);
    isRunning = true;
    updateControls();
    timer = setInterval(() => {
        if (currentSeconds > 0) {
            currentSeconds--;
            updateDisplay();
        } else {
            clearInterval(timer);
            isRunning = false;
            currentSeconds = 0;
            updateDisplay();
            onComplete();
        }
    }, 1000);
}

function pauseTimer() {
    clearInterval(timer);
    isRunning = false;
    updateControls();
}

function resetTimer() {
    clearInterval(timer);
    isRunning = false;
    currentSeconds = defaultTime;
    updateDisplay();
}

function getAudioContext() {
    if (!audioCtx) {
        const AudioContext =
            window.AudioContext || window.webkitAudioContext;
        if (!AudioContext) return null;
        audioCtx = new AudioContext();
    }
    if (audioCtx.state === "suspended") {
        audioCtx.resume();
    }
    return audioCtx;
}

function loadAlertSound(ctx) {
    if (alertSoundBuffer) return Promise.resolve(alertSoundBuffer);
    return fetch("audio/achievement.ogg")
        .then((res) => res.arrayBuffer())
        .then((data) => ctx.decodeAudioData(data))
        .then((buffer) => {
            alertSoundBuffer = buffer;
            return buffer;
        })
        .catch(() => null);
}

function playAlertSound() {
    const ctx = getAudioContext();
    if (!ctx) return;

    const play = () => {
        loadAlertSound(ctx).then((buffer) => {
            if (buffer) {
                const source = ctx.createBufferSource();
                source.buffer = buffer;
                const gain = ctx.createGain();
                gain.gain.value = 1;
                source.connect(gain);
                gain.connect(ctx.destination);
                source.start();
            } else {
                playBeep(ctx);
            }
        });
    };

    if (ctx.state === "suspended") {
        ctx.resume().then(play).catch(play);
    } else {
        play();
    }
}

function playBeep(ctx) {
    const now = ctx.currentTime;
    for (let i = 0; i < 3; i++) {
        const offset = now + i * 0.5;
        const osc = ctx.createOscillator();
        const gain = ctx.createGain();
        osc.type = "sine";
        osc.frequency.value = 880;
        gain.gain.setValueAtTime(0.0001, offset);
        gain.gain.exponentialRampToValueAtTime(0.6, offset + 0.03);
        gain.gain.exponentialRampToValueAtTime(0.0001, offset + 0.38);
        osc.connect(gain);
        gain.connect(ctx.destination);
        osc.start(offset);
        osc.stop(offset + 0.42);
    }
}

function onComplete() {
    playAlertSound();
    const title = activeAlert ? activeAlert.name : t("attention");
    const message = t("timeUpMessage");
    openDialog(title, message);
    if (sessionId) {
        fetch("/notifier/notify", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                sessionId,
                title,
                body: message,
            }),
        })
            .then(() => console.log("Notification sent successfully"))
            .catch((e) => console.log("Notification failed", e));
    }
}

function openDialog(title, message) {
    dialogTitle.textContent = title;
    dialogMessage.textContent = message;
    dialog.classList.add("show");
}

function closeDialog() {
    dialog.classList.remove("show");
}

function selectAlert(alert) {
    activeAlert = alert;
    defaultTime = totalOf(alert);
    totalSeconds = defaultTime;
    pageTitle.textContent = alert.name;
    resetTimer();
    renderAlerts();
}

function renderAlerts() {
    alertsList.innerHTML = "";
    const hasAlerts = alerts.length > 0;
    emptyState.hidden = hasAlerts;
    alertsCount.textContent = String(alerts.length);

    alerts.forEach((al) => {
        const li = document.createElement("li");
        if (activeAlert && activeAlert.id === al.id) {
            li.classList.add("active");
        }

        const info = document.createElement("div");
        info.className = "alert-info";

        const nameEl = document.createElement("span");
        nameEl.className = "alert-name";
        nameEl.textContent = al.name;

        const timeEl = document.createElement("span");
        timeEl.className = "alert-time";
        timeEl.textContent = formatTime(totalOf(al));

        info.appendChild(nameEl);
        info.appendChild(timeEl);

        const actions = document.createElement("div");
        actions.className = "actions";

        const useBtn = document.createElement("button");
        useBtn.className = "btn-use";
        useBtn.type = "button";
        useBtn.innerHTML = `${ICONS.use}<span>${t("use")}</span>`;
        useBtn.onclick = () => selectAlert(al);

        const editBtn = document.createElement("button");
        editBtn.className = "btn-edit";
        editBtn.type = "button";
        editBtn.innerHTML = `${ICONS.edit}<span>${t("edit")}</span>`;
        editBtn.onclick = () => editAlert(al);

        const deleteBtn = document.createElement("button");
        deleteBtn.className = "btn-delete";
        deleteBtn.type = "button";
        deleteBtn.innerHTML = `${ICONS.delete}<span>${t("delete")}</span>`;
        deleteBtn.onclick = () => removeAlert(al);

        actions.appendChild(useBtn);
        actions.appendChild(editBtn);
        actions.appendChild(deleteBtn);

        li.appendChild(info);
        li.appendChild(actions);
        alertsList.appendChild(li);
    });
}

function removeAlert(alert) {
    const index = alerts.indexOf(alert);
    const [removed] = alerts.splice(index, 1);

    if (activeAlert && activeAlert.id === removed.id) {
        activeAlert = alerts[0] || null;
        if (activeAlert) {
            selectAlert(activeAlert);
        } else {
            defaultTime = 0;
            totalSeconds = 0;
            pageTitle.textContent = document.title;
            resetTimer();
        }
    }

    renderAlerts();
    showSnackbar(`${t("deleted")} "${removed.name}"`, t("undo"), () => {
        alerts.splice(index, 0, removed);
        renderAlerts();
    });
}

let snackbarTimeout;
function showSnackbar(message, actionLabel, onAction, duration = 4000) {
    snackbar.innerHTML = "";
    const messageEl = document.createElement("span");
    messageEl.textContent = message;
    snackbar.appendChild(messageEl);

    if (actionLabel && onAction) {
        const actionBtn = document.createElement("button");
        actionBtn.className = "snackbar-action";
        actionBtn.textContent = actionLabel;
        actionBtn.onclick = () => {
            clearTimeout(snackbarTimeout);
            snackbar.classList.remove("show");
            onAction();
        };
        snackbar.appendChild(actionBtn);
        duration = Math.max(duration, 10000);
    }

    snackbar.classList.add("show");
    clearTimeout(snackbarTimeout);
    snackbarTimeout = setTimeout(
        () => snackbar.classList.remove("show"),
        duration,
    );
}

function clearForm() {
    editingId = null;
    nameInput.value = "";
    hoursInput.value = 0;
    minutesInput.value = 0;
    secondsInput.value = 0;
    addAlertLabel.textContent = t("addAlert");
    cancelFormLabel.textContent = t("cancel");
    nameInput.focus();
}

function editAlert(alert) {
    editingId = alert.id;
    nameInput.value = alert.name;
    hoursInput.value = alert.hours;
    minutesInput.value = alert.minutes;
    secondsInput.value = alert.seconds;
    addAlertLabel.textContent = t("save");
    cancelFormLabel.textContent = t("cancelEdit");
    nameInput.focus();
    alertForm.scrollIntoView({ behavior: "smooth", block: "center" });
}

alertForm.addEventListener("submit", (e) => {
    e.preventDefault();
    const name = nameInput.value.trim();
    const hours = parseInt(hoursInput.value) || 0;
    const minutes = parseInt(minutesInput.value) || 0;
    const seconds = parseInt(secondsInput.value) || 0;
    const total = hours * 3600 + minutes * 60 + seconds;

    if (!name || total <= 0) {
        showSnackbar(t("invalidAlert"));
        if (!name) nameInput.focus();
        return;
    }

    if (editingId !== null) {
        const existing = alerts.find((a) => a.id === editingId);
        if (existing) {
            existing.name = name;
            existing.hours = hours;
            existing.minutes = minutes;
            existing.seconds = seconds;
            clearForm();
            if (activeAlert && activeAlert.id === existing.id) {
                selectAlert(existing);
            } else {
                renderAlerts();
            }
            showSnackbar(`${t("updated")} "${name}"`);
        }
        return;
    }

    const alert = { id: Date.now(), name, hours, minutes, seconds };
    alerts.push(alert);
    clearForm();
    selectAlert(alert);
    showSnackbar(`${t("added")} "${name}"`);
});

cancelFormBtn.addEventListener("click", clearForm);
[hoursInput, minutesInput, secondsInput].forEach((input) => {
    input.addEventListener("focus", () => {
        if (input.value === "0") input.value = "";
    });
    input.addEventListener("blur", () => {
        if (input.value === "") input.value = "0";
    });
});
startBtn.addEventListener("click", startTimer);
pauseBtn.addEventListener("click", pauseTimer);
resetBtn.addEventListener("click", resetTimer);
dialogOk.addEventListener("click", closeDialog);

dialogSnooze.addEventListener("click", () => {
    closeDialog();
    currentSeconds += SNOOZE_SECONDS;
    totalSeconds = currentSeconds;
    updateDisplay();
    startTimer();
    showSnackbar(t("snoozed"));
});

document.addEventListener("keydown", (e) => {
    const tag = e.target.tagName;
    if (e.code === "Space" && tag !== "INPUT" && tag !== "BUTTON" && tag !== "TEXTAREA") {
        e.preventDefault();
        if (dialog.classList.contains("show")) return;
        isRunning ? pauseTimer() : startTimer();
    }
});

function updateCurrentTime() {
    const date = new Date();
    const h = date.getHours().toString().padStart(2, "0");
    const m = date.getMinutes().toString().padStart(2, "0");
    const s = date.getSeconds().toString().padStart(2, "0");
    currentTimeDisplay.textContent = `${h}:${m}:${s}`;
}

language = detectLanguage();
applyLanguage();
if (alerts.length > 0) {
    selectAlert(alerts[0]);
} else {
    resetTimer();
}
renderAlerts();
setInterval(updateCurrentTime, 1000);
updateCurrentTime();
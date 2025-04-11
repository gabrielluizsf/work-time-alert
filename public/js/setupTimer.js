function setupTimer(language) {
    const prefix = language === "pt" ? "" : "-en";
    const workTimeInput = document.getElementById(`work-time${prefix}`);
    const startBtn = document.getElementById(`start-btn${prefix}`);
    const stopBtn = document.getElementById(`stop-btn${prefix}`);
    const timerDisplay = document.getElementById(`timer${prefix}`);
    const alertDiv = document.getElementById(`alert${prefix}`);
    const currentTimeDisplay = document.getElementById(`current-time${prefix}`);

    let countdownInterval = null;
    let targetTime = null;

    setInterval(() => updateCurrentTime(currentTimeDisplay, language), 1000);
    updateCurrentTime(currentTimeDisplay, language);

    startBtn.addEventListener("click", () => startCountdown(language));
    stopBtn.addEventListener("click", () => stopCountdown(language));

    function startCountdown(lang) {
        if (!workTimeInput.value) {
            alert(
                lang === "pt"
                    ? "Por favor, defina um horário válido"
                    : "Please set a valid time",
            );
            return;
        }

        const [hours, minutes] = workTimeInput.value.split(":").map(Number);
        const now = new Date();
        const workTime = hours + minutes;
        targetTime = now;
        const nowTime = now.getHours() + now.getMinutes();
        if (workTime < nowTime || workTime == nowTime) {
            alert(
                lang === "pt"
                    ? "Por favor, defina um horário válido"
                    : "Please set a valid time",
            );
            return;
        }
        targetTime.setHours(hours, minutes, 0, 0);

        startBtn.disabled = true;
        stopBtn.disabled = false;
        workTimeInput.disabled = true;
        alertDiv.style.display = "none";

        updateCountdown(lang);
        countdownInterval = setInterval(() => updateCountdown(lang), 1000);
    }

    function stopCountdown(lang) {
        clearInterval(countdownInterval);
        countdownInterval = null;
        timerDisplay.textContent = lang === "pt"
            ? "Contador parado. Defina um novo horário."
            : "Timer stopped. Set a new time.";
        stopBtn.textContent = lang === "pt" ? "Parar Contador" : "Stop Timer";
        startBtn.disabled = false;
        stopBtn.disabled = true;
        workTimeInput.disabled = false;
        alertDiv.style.display = "none";
        workTimeInput.value = "";
    }

    function updateCountdown(lang) {
        const now = new Date();
        const diff = targetTime - now;

        if (diff <= 0) {
            timerDisplay.textContent = lang === "pt"
                ? "Chegou a hora!"
                : "Time's up!";
            stopBtn.textContent = lang === "pt" ? "Reiniciar Contador" : "Restart Timer";
            alertDiv.style.display = "block";
            clearInterval(countdownInterval);
            playAlertSound();
            return;
        }

        const hours = Math.floor(diff / (1000 * 60 * 60));
        const minutes = Math.floor((diff % (1000 * 60 * 60)) / (1000 * 60));
        const seconds = Math.floor((diff % (1000 * 60)) / 1000);

        timerDisplay.textContent = lang === "pt"
            ? `Tempo restante: ${hours}h ${minutes}m ${seconds}s`
            : `Time remaining: ${hours}h ${minutes}m ${seconds}s`;
    }

    function updateCurrentTime(element, lang) {
        const now = new Date();
        const hours = now.getHours().toString().padStart(2, "0");
        const minutes = now.getMinutes().toString().padStart(2, "0");
        const seconds = now.getSeconds().toString().padStart(2, "0");
        element.textContent = lang === "pt"
            ? `Horário atual: ${hours}:${minutes}:${seconds}`
            : `Current time: ${hours}:${minutes}:${seconds}`;
    }

    function playAlertSound() {
        try {
            const audioContext =
                new (window.AudioContext || window.webkitAudioContext)();
            const oscillator = audioContext.createOscillator();
            const gainNode = audioContext.createGain();

            oscillator.type = "sine";
            oscillator.frequency.value = 880;
            gainNode.gain.value = 0.5;

            oscillator.connect(gainNode);
            gainNode.connect(audioContext.destination);

            oscillator.start();

            setTimeout(() => {
                oscillator.stop();
            }, 1000);
        } catch (e) {
            console.log("Alert sound error:", e);
        }
    }
}

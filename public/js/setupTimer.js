function setupTimer(language) {
    const defaultLanguage = "pt"
    const prefix = language === defaultLanguage ? "" : "-en";
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
    const clickEvent = "click"
    startBtn.addEventListener(clickEvent, () => startCountdown(language));
    stopBtn.addEventListener(clickEvent, () => stopCountdown(language));

    function startCountdown(lang) {
        const alertInvalidTime = ()=> alert(
            lang === defaultLanguage
                ? "Por favor, defina um horário válido"
                : "Please set a valid time",
        );
        if (!workTimeInput.value) {
            alertInvalidTime();
            return;
        }

        const [hours, minutes] = workTimeInput.value.split(":").map(Number);
        targetTime = now();

        const chosenTime = now();
        chosenTime.setHours(hours, minutes, 0, 0);
        
        if (chosenTime <= now()) {
            alertInvalidTime();
            return;
        }
    
        targetTime = chosenTime;

        startBtn.disabled = true;
        stopBtn.disabled = false;
        workTimeInput.disabled = true;
        alertDiv.style.display = "none";

        updateCountdown(lang);
        countdownInterval = setInterval(() => updateCountdown(lang), 1000);
    }

    function now() {
        return new Date();
    }

    function stopCountdown(lang) {
        clearInterval(countdownInterval);
        countdownInterval = null;
        timerDisplay.textContent = lang === defaultLanguage
            ? "Contador parado. Defina um novo horário."
            : "Timer stopped. Set a new time.";
        stopBtn.textContent = lang === defaultLanguage ? "Parar Contador" : "Stop Timer";
        startBtn.disabled = false;
        stopBtn.disabled = true;
        workTimeInput.disabled = false;
        alertDiv.style.display = "none";
        workTimeInput.value = "";
    }

    function updateCountdown(lang) {
        const second = 1000;
        const minute = second * 60;
        const hour = minute * 60;
        const diff = targetTime - now();

        if (diff <= 0) {
            timerDisplay.textContent = lang === defaultLanguage
                ? "Chegou a hora!"
                : "Time's up!";
            stopBtn.textContent = lang === defaultLanguage ? "Reiniciar Contador" : "Restart Timer";
            alertDiv.style.display = "block";
            clearInterval(countdownInterval);
            playAlertSound();
            return;
        }

        const hours = Math.floor(diff / hour);
        const minutes = Math.floor((diff % hour) / minute);
        const seconds = Math.floor((diff % hour) / second);

        timerDisplay.textContent = lang === defaultLanguage
            ? `Tempo restante: ${hours}h ${minutes}m ${seconds}s`
            : `Time remaining: ${hours}h ${minutes}m ${seconds}s`;
    }

    function updateCurrentTime(element, lang) {
        const date = now();
        const hours = date.getHours().toString().padStart(2, "0");
        const minutes = date.getMinutes().toString().padStart(2, "0");
        const seconds = date.getSeconds().toString().padStart(2, "0");
        element.textContent = lang === defaultLanguage
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

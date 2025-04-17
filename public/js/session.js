async function initSession(sessions) {
    if (sessions.length > 0) {
        return;
    }
    const res = await fetch("/notifier/session", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
    });
    const data = await res.json();
    sessions.push(data.sessionId);
}

let cachedSubscriptions = new Map();

async function registerServiceWorker(sessionId) {
    try {
        await navigator.serviceWorker.register(
            "/service-worker.js",
        );
        const registration = await navigator
            .serviceWorker.ready;
        const exists = await registration
            .pushManager.getSubscription();
        if (!exists) {
            const res = await fetch(
                "/notifier/public",
            );
            const { key: publicKey } = await res.json();
            const subscription = await registration
                .pushManager.subscribe({
                    userVisibleOnly: true,
                    applicationServerKey: urlBase64ToUint8Array(
                        publicKey,
                    ),
                });

            await fetch(
                "/notifier/subscribe",
                {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                    },
                    body: JSON.stringify({
                        sessionId,
                        subscription,
                    }),
                },
            );
        }
    } catch (error) {
        console.log("erro ao registrar service worker", error);
    }
}

const urlBase64ToUint8Array = (publicKey) => {
    const padding = "=".repeat((4 - (publicKey.length % 4)) % 4);
    const base64 = (publicKey + padding)
        .replace(/-/g, "+")
        .replace(/_/g, "/");

    try {
        const rawData = window.atob(base64);
        const outputArray = new Uint8Array(rawData.length);
        for (let i = 0; i < rawData.length; i += 1) {
            outputArray[i] = rawData.charCodeAt(i);
        }
        return outputArray;
    } catch (error) {
        throw new Error(
            `Failed to convert base64 to Uint8Array: ${error.message}`,
        );
    }
};

self.addEventListener('push', async event => {
    const data = await event.data.json();
    if (Notification.permission === 'granted') {
        const {body, title} = data??{};

        event.waitUntil(
            self.registration.showNotification(title, {
                body: body,
            })
        );
    } else {
        console.log('Notification permission not granted');
    }
});
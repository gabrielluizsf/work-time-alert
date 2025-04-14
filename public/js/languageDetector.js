function detectLanguage() {
    const userLanguage = navigator.language || navigator.userLanguage;
    const isBrazil = userLanguage.toLowerCase() === 'pt-br' || 
                     (navigator.geolocation && confirm("Você está no Brasil? (Are you in Brazil?)"));
    const brazilContent =  document.getElementById('pt-br-content') 
    const usaContent = document.getElementById('en-us-content');
    if (isBrazil) {
        document.documentElement.lang = 'pt-BR';
        brazilContent.style.display = 'block';
        usaContent.style.display = 'none';
        return 'pt';
    } else {
        document.documentElement.lang = 'en-US';
        brazilContent.style.display = 'none';
        usaContent.style.display = 'block';
        return 'en';
    }
}
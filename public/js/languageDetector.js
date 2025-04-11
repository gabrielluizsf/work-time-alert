function detectLanguage() {
    const userLanguage = navigator.language || navigator.userLanguage;
    const isBrazil = userLanguage.toLowerCase() === 'pt-br' || 
                     (navigator.geolocation && confirm("Você está no Brasil? (Are you in Brazil?)"));
    
    if (isBrazil) {
        document.documentElement.lang = 'pt-BR';
        document.getElementById('pt-br-content').style.display = 'block';
        document.getElementById('en-us-content').style.display = 'none';
        return 'pt';
    } else {
        document.documentElement.lang = 'en-US';
        document.getElementById('pt-br-content').style.display = 'none';
        document.getElementById('en-us-content').style.display = 'block';
        return 'en';
    }
}
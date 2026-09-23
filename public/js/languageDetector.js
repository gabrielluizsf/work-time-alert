function detectLanguage() {
    const userLanguage = navigator.language || navigator.userLanguage;
    const isBrazil = (userLanguage || "").toLowerCase().startsWith("pt");
    document.documentElement.lang = isBrazil ? "pt-BR" : "en-US";
    return isBrazil ? "pt" : "en";
}
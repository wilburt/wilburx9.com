// Script for processing html in default.hbs


!function () {
    handleDarkMode()
    setAboutHref()
}();

function handleDarkMode() {
    let modeQueryList = window.matchMedia('(prefers-color-scheme: dark)');
    const isDarkMode = () => localStorage.theme === 'dark' || (!('theme' in localStorage) && modeQueryList.matches);

    if (isDarkMode()) {
        document.documentElement.classList.add('dark')
    } else {
        document.documentElement.classList.remove('dark')
    }

    if (!('theme' in localStorage)) {
        modeQueryList.addEventListener('change', function () {
            if (!('theme' in localStorage)) changeMode(true, isDarkMode())
        });
    }

    document.addEventListener("DOMContentLoaded", () => {
        document.getElementById("theme-switch").addEventListener("click", () => {
            changeMode(false, !isDarkMode())
        });
    });

    function changeMode(fromSystem, darkMode) {
        if (darkMode) {
            if (!fromSystem) localStorage.theme = 'dark'
            document.documentElement.classList.add('dark')
        } else {
            if (!fromSystem) localStorage.theme = 'light'
            document.documentElement.classList.remove('dark')
        }
    }
}

function setAboutHref() {
    document.addEventListener("DOMContentLoaded", () => {
        document.querySelectorAll('[id=about-me]').forEach(e => e.href = window.origin)
    });
}
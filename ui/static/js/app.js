const toggleTheme = () => {
  const root = document.documentElement;
  const currentTheme = root.getAttribute("data-theme");

  if (currentTheme) {
    const next = currentTheme === "light" ? "dark" : "light";
    root.setAttribute("data-theme", next);
    localStorage.setItem("theme", next);
    return;
  }

  const prefersDark = window.matchMedia("(prefers-color-scheme: dark)").matches;
  const next = prefersDark ? "light" : "dark";
  root.setAttribute("data-theme", next);
  localStorage.setItem("theme", next);
};

const changeColorScheme = (e) => {
  const value = e.target.value;
  const root = document.querySelector(":root");
  root.style.setProperty("--primary-hue", value);
};

const themeToggle = document.querySelector("#theme-toggle");
const mobileThemeToggle = document.querySelector("#mobile-theme-toggle");
themeToggle.addEventListener("click", toggleTheme);
mobileThemeToggle.addEventListener("click", toggleTheme);

const colorTheme = document.querySelector("#color-theme");
const mobileColortheme = document.querySelector("#mobile-color-theme");
colorTheme.addEventListener("change", changeColorScheme);
mobileColortheme.addEventListener("change", changeColorScheme);

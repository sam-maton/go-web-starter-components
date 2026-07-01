const themeToggle = document.querySelector("#theme-toggle");
themeToggle.addEventListener("click", () => {
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
});

const colorTheme = document.querySelector("#color-theme");
colorTheme.addEventListener("change", () => {
  const value = colorTheme.value;
  const root = document.querySelector(":root");
  root.style.setProperty("--primary-hue", value);
});

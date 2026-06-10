const themeToggle = document.querySelector("#theme-toggle");
themeToggle.addEventListener("click", () => {
  const root = document.documentElement;
  const currentTheme = root.getAttribute("data-theme");
  let next = null;
  if (currentTheme) {
    next = currentTheme === "light" ? "dark" : "light";
    root.setAttribute("data-theme", next);
    localStorage.setItem("theme", next);
    return;
  }

  const prefersDark = window.matchMedia("(prefers-color-scheme: dark)").matches;
  next = prefersDark ? "light" : "dark";
  root.setAttribute("data-theme", next);
  localStorage.setItem("theme", next);
});

// DROPDOWN MENU
const anchorElement = document.getElementById("tooltip");
const anchor = document.getElementById("anchor");

anchor.addEventListener("mouseover", () => {
  anchorElement.classList.add("show");
});

anchor.addEventListener("mouseout", () => {
  anchorElement.classList.remove("show");
});

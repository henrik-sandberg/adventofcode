document.addEventListener("DOMContentLoaded", () => {
  const mainTitle = document.getElementById("main-title");
  const yearsListDiv = document.getElementById("years-list");
  const yearsButtonsDiv = document.getElementById("years-buttons");
  const daysListDiv = document.getElementById("days-list");
  const daysButtonsDiv = document.getElementById("days-buttons");
  const daysTitle = document.getElementById("days-title");
  const problemSolverDiv = document.getElementById("problem-solver");
  const problemTitle = document.getElementById("problem-title");
  const puzzleInputTextarea = document.getElementById("puzzle-input");
  const fileInput = document.getElementById("file-input");
  const inputError = document.getElementById("input-error");
  const solveButton = document.getElementById("solve-button");
  const solutionOutput = document.getElementById("solution-output");
  const backToYearsButton = document.getElementById("back-to-years");
  const backToDaysButton = document.getElementById("back-to-days");

  const views = document.querySelectorAll(".view");
  let currentYear = null;
  let currentDay = null;

  // --- Navigation Functions ---

  function showView(viewToShow) {
    views.forEach((view) => {
      view.classList.remove("active");
    });
    viewToShow.classList.add("active");
  }

  function showYearsList() {
    mainTitle.classList.remove("hidden");
    showView(yearsListDiv);
    solutionOutput.textContent = ""; // Clear previous output
    puzzleInputTextarea.value = ""; // Clear text input
    fileInput.value = ""; // Clear file input
  }

  function showDaysList(year) {
    mainTitle.classList.add("hidden");
    currentYear = year;
    showView(daysListDiv);
    daysTitle.textContent = `Select a Day for ${year}`;
    solutionOutput.textContent = ""; // Clear previous output
    puzzleInputTextarea.value = ""; // Clear text input
    fileInput.value = ""; // Clear file input
  }

  function showProblemSolver(year, day) {
    mainTitle.classList.add("hidden");
    currentYear = year;
    currentDay = day;
    showView(problemSolverDiv);
    problemTitle.textContent = `Solving ${year} - Day ${day}`;
    solutionOutput.textContent = ""; // Clear previous output
    inputError.textContent = ""; // Clear input error
  }

  // --- Fetch Data from Backend ---

  async function fetchYears() {
    try {
      const response = await fetch("/api/years");
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const years = await response.json();
      yearsButtonsDiv.innerHTML = "";
      years.forEach((year) => {
        const button = document.createElement("button");
        button.textContent = year;
        button.onclick = () => fetchDays(year);
        yearsButtonsDiv.appendChild(button);
      });
    } catch (error) {
      console.error("Error fetching years:", error);
      yearsButtonsDiv.innerHTML = "<p>Error loading years.</p>";
    }
  }

  async function fetchDays(year) {
    showDaysList(year);
    try {
      const response = await fetch(`/api/years/${year}/days`);
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const implementedDays = new Set(await response.json());
      daysButtonsDiv.innerHTML = "";

      for (let i = 1; i <= 25; i++) {
        const day = String(i).padStart(2, "0");
        const button = document.createElement("button");
        button.textContent = `Day ${day}`;

        if (implementedDays.has(day)) {
          button.onclick = () => showProblemSolver(year, day);
        } else {
          button.disabled = true;
          const tooltip = document.createElement("span");
          tooltip.className = "tooltip";
          tooltip.textContent = "This day has not been implemented yet.";
          button.appendChild(tooltip);
        }
        daysButtonsDiv.appendChild(button);
      }
    } catch (error) {
      console.error(`Error fetching days for ${year}:`, error);
      daysButtonsDiv.innerHTML = `<p>Error loading days for ${year}.</p>`;
    }
  }

  async function solvePuzzle() {
    const textInput = puzzleInputTextarea.value;
    const file = fileInput.files[0];
    inputError.textContent = "";

    if (textInput && file) {
      inputError.textContent =
        "Please provide input either via text field OR file, not both.";
      return;
    }

    let inputData = "";
    if (file) {
      try {
        inputData = await file.text();
      } catch (error) {
        inputError.textContent = "Error reading file.";
        console.error("Error reading file:", error);
        return;
      }
    } else if (textInput) {
      inputData = textInput;
    } else {
      inputError.textContent = "Please provide puzzle input.";
      return;
    }

    try {
      const response = await fetch(
        `/api/years/${currentYear}/days/${currentDay}`,
        {
          method: "POST",
          headers: {
            "Content-Type": "text/plain",
          },
          body: inputData,
        },
      );

      const result = await response.text();
      if (!response.ok) {
        solutionOutput.textContent = `Error: ${result}`;
        solutionOutput.style.color = "red";
      } else {
        solutionOutput.textContent = result;
        solutionOutput.style.color = "white";
      }
    } catch (error) {
      console.error("Error solving puzzle:", error);
      solutionOutput.textContent = `An unexpected error occurred: ${error.message}`;
      solutionOutput.style.color = "red";
    }
  }

  // --- Event Listeners ---

  backToYearsButton.addEventListener("click", showYearsList);
  backToDaysButton.addEventListener("click", () => fetchDays(currentYear));
  solveButton.addEventListener("click", solvePuzzle);

  // Initial load
  fetchYears();
  showYearsList();
});

document.addEventListener("DOMContentLoaded", () => {
  const AppState = {
    year: null,
    day: null,
  };

  const mainTitle = document.getElementById("main-title");
  const yearsListDiv = document.getElementById("years-list");
  const yearsButtonsDiv = document.getElementById("years-buttons");
  const daysListDiv = document.getElementById("days-list");
  const daysButtonsDiv = document.getElementById("days-buttons");
  const daysTitle = document.getElementById("days-title");
  const problemSolverDiv = document.getElementById("problem-solver");
  const problemTitle = document.getElementById("problem-title");

  const puzzleInputTextarea = document.getElementById("puzzle-input");
  const inputError = document.getElementById("input-error");

  const solutionOutput = document.getElementById("solution-output");

  const solveButton = document.getElementById("solve-button");
  const solveButtonText = solveButton.querySelector(".button-text");
  const solveButtonSpinner = solveButton.querySelector(".spinner");
  const backToYearsButton = document.getElementById("back-to-years");
  const backToDaysButton = document.getElementById("back-to-days");

  const views = document.querySelectorAll(".view");

  function showView(viewToShow) {
    views.forEach((view) => view.classList.remove("active"));
    viewToShow.classList.add("active");
  }

  function showYearsList() {
    mainTitle.classList.remove("hidden");
    showView(yearsListDiv);
  }

  function showDaysList() {
    mainTitle.classList.add("hidden");
    showView(daysListDiv);
    daysTitle.textContent = `Select a Day for ${AppState.year}`;
    solutionOutput.textContent = "";
    puzzleInputTextarea.value = "";
  }

  function showProblemSolver() {
    mainTitle.classList.add("hidden");
    showView(problemSolverDiv);
    problemTitle.textContent = `Solving ${AppState.year} - Day ${AppState.day}`;
    solutionOutput.textContent = "";
    inputError.textContent = "";
  }

  async function fetchYears() {
    try {
      const response = await fetch("/api/years");
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      yearsButtonsDiv.innerHTML = "";
      const years = (await response.json()).data;
      years.forEach((year) => {
        const button = document.createElement("button");
        button.textContent = year;
        button.onclick = () => {
          AppState.year = year;
          fetchDays();
        };
        yearsButtonsDiv.appendChild(button);
      });
    } catch (error) {
      console.error("Error fetching years:", error);
      yearsButtonsDiv.innerHTML = "<p>Error loading years.</p>";
    }
  }

  async function fetchDays() {
    showDaysList();
    try {
      const response = await fetch(`/api/years/${AppState.year}/days`);
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const implementedDays = new Set((await response.json()).data);
      daysButtonsDiv.innerHTML = "";

      for (let i = 1; i <= 25; i++) {
        const day = String(i).padStart(2, "0");
        const button = document.createElement("button");
        button.textContent = `Day ${day}`;

        if (implementedDays.has(day)) {
          button.onclick = () => {
            AppState.day = day;
            showProblemSolver();
          };
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
      console.error(`Error fetching days for ${AppState.year}:`, error);
      daysButtonsDiv.innerHTML = `<p>Error loading days for ${AppState.year}.</p>`;
    }
  }

  async function solvePuzzle() {
    solveButton.disabled = true;
    solveButtonText.classList.add("hidden");
    solveButtonSpinner.classList.remove("hidden");

    const textInput = puzzleInputTextarea.value;
    inputError.textContent = "";

    try {
      if (!textInput) {
        inputError.textContent = "Please provide puzzle input.";
        return;
      }
      const response = await fetch(
        `/api/years/${AppState.year}/days/${AppState.day}`,
        {
          method: "POST",
          headers: {
            "Content-Type": "text/plain",
          },
          body: textInput,
        },
      );
      const result = await response.text();
      if (response.ok) {
        solutionOutput.textContent = result;
        solutionOutput.style.color = "white";
      } else {
        solutionOutput.textContent = `Error: ${result}`;
        solutionOutput.style.color = "red";
      }
    } catch (error) {
      solutionOutput.textContent = `An unexpected error occurred: ${error.message}`;
      solutionOutput.style.color = "red";
    } finally {
      solveButton.disabled = false;
      solveButtonText.classList.remove("hidden");
      solveButtonSpinner.classList.add("hidden");
    }
  }

  backToYearsButton.addEventListener("click", showYearsList);
  backToDaysButton.addEventListener("click", fetchDays);
  solveButton.addEventListener("click", solvePuzzle);

  fetchYears();
  showYearsList();
});

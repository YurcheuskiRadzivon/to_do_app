document.addEventListener("DOMContentLoaded", function () {
  tasks.forEach((task) => addTask(task.name, task.description));

  listOfTasks.addEventListener("click", function (event) {
    const target = event.target.closest("button, label"); // Найти ближайшую кнопку или label

    if (target && target.classList.contains("action_button")) {
      const actionsPanel = target.nextElementSibling;
      actionsPanel.style.display =
        actionsPanel.style.display === "none" ? "block" : "none";
      target.style.display = "none"; // Скрыть кнопку действий
    }

    if (target && target.htmlFor === "delete") {
      const taskElement = target.closest(".task_line");
      taskElement.remove();
      
    }

    if (target && target.htmlFor === "done") {
      const taskElement = target.closest(".task_line");
      taskElement.querySelector(".content").style.textDecoration =
        "line-through";
    }

    if (target && target.htmlFor === "rename") {
      const taskElement = target.closest(".task_line");
      const newName = prompt("Enter new task name:");
      const newDesc = prompt("Enter new task description:");
      if (newName && newDesc) {
        taskElement.querySelector(".heading").textContent = newName;
        taskElement.querySelector(".para").textContent = newDesc;
      }
    }
  });
});

const submitBtn = document.getElementById("submit_task");
const listOfTasks = document.getElementById("list__tasks");
const taskName = document.getElementById("name");
const taskDescript = document.getElementById("description");
var ind = 0;

var tasks = JSON.parse(localStorage.getItem("tasks")) || [];

function saveTasks() {
  localStorage.setItem("tasks", JSON.stringify(tasks));
}

function addTask(name, desc) {
  listOfTasks.insertAdjacentHTML(
    "beforeend",
    `
  <li class="task_line">
    <div class="task" data-ind="${ind++}">
      <div class="content">
        <p class="heading">${name}</p>
        <p class="para">${desc}</p>
        <p class="para para-sm">Jan 1, 2024</p>
      </div>
    </div>
    <button class="action_button">
      <svg viewBox="0 0 256 256" xmlns="http://www.w3.org/2000/svg">
      <path d="M140,128a12,12,0,1,1-12-12A12,12,0,0,1,140,128ZM128,72a12,12,0,1,0-12-12A12,12,0,0,0,128,72Zm0,112a12,12,0,1,0,12,12A12,12,0,0,0,128,184Z"></path></svg>
    </button>
    <div class="user-actions-with-task" style="display: none">
      <div class="card">
          <ul
            class="list"
            style="--color:#5353ff;--hover-storke:#fff; --hover-color:#fff"
          >
            <li class="element">
              <label for="rename">
                <input type="radio" id="rename" name="filed" checked="" />
                <svg
                  class="lucide lucide-pencil"
                  stroke-linejoin="round"
                  stroke-linecap="round"
                  stroke-width="2"
                  stroke="#7e8590"
                  fill="none"
                  viewBox="0 0 24 24"
                  height="25"
                  width="25"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path
                    d="M21.174 6.812a1 1 0 0 0-3.986-3.987L3.842 16.174a2 2 0 0 0-.5.83l-1.321 4.352a.5.5 0 0 0 .623.622l4.353-1.32a2 2 0 0 0 .83-.497z"
                  ></path>
                  <path d="m15 5 4 4"></path>
                </svg>
                Rename</label
              >
            </li>
         
            <div class="separator"></div>
            <li class="element" style="--color:#5353ff">
              <label for="settings">
                <input type="radio" id="done" name="filed" />
                <svg
                xmlns="http://www.w3.org/2000/svg"
                x="0px"
                y="0px"
                width="40"
                height="40"
                viewBox="0 0 24 24"
                fill="none"
                  stroke="#7e8590"
                  stroke-width="2"
                  stroke-linecap="round"
                  stroke-linejoin="round"
              >
                <path
                  fill="none" d="M 20.292969 5.2929688 L 9 16.585938 L 4.7070312 12.292969 L 3.2929688 13.707031 L 9 19.414062 L 21.707031 6.7070312 L 20.292969 5.2929688 z"
                ></path>
              </svg>
                Done</label
              >
            </li>
            <li class="element delete" style="--color:#8e2a2a">
              <label for="delete">
                <input type="radio" id="delete" name="filed" />
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  width="24"
                  height="24"
                  viewBox="0 0 24 24"
                  fill="none"
                  stroke="#7e8590"
                  stroke-width="2"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  class="lucide lucide-trash-2"
                >
                  <path d="M3 6h18"></path>
                  <path d="M19 6v14c0 1-1 2-2 2H7c-1 0-2-1-2-2V6"></path>
                  <path d="M8 6V4c0-1 1-2 2-2h4c1 0 2 1 2 2v2"></path>
                  <line x1="10" x2="10" y1="11" y2="17"></line>
                  <line x1="14" x2="14" y1="11" y2="17"></line>
                </svg>
                Delete</label
              >
            </li>
            <div class="separator"></div>
            <!-- <li
              class="element"
              style="--color:rgba(56, 45, 71, 0.836);--hover-storke:#bd89ff;--hover-color:#bd89ff"
            >
              <label for="teamaccess">
                <input type="radio" id="teamaccess" name="filed" />
  
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  width="24"
                  height="24"
                  viewBox="0 0 24 24"
                  fill="none"
                  stroke="#7e8590"
                  stroke-width="2"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  class="lucide lucide-users-round"
                >
                  <path d="M18 21a8 8 0 0 0-16 0"></path>
                  <circle cx="10" cy="8" r="5"></circle>
                  <path d="M22 20c0-3.37-2-6.5-4-8a5 5 0 0 0-.45-8.3"></path>
                </svg>
                Team Access</label
              >
            </li> -->
          </ul>
        </div>
    </div>
  </li>
  `
  );
  taskDescript.value = "";
  taskName.value = "";
}

submitBtn.onclick = function () {
  if (taskName.value.length != 0 && taskDescript.value.length != 0) {
    var newTask = { name: taskName.value, description: taskDescript.value };
    tasks.push(newTask);
    addTask(taskName.value, taskDescript.value);
    saveTasks();
  } else {
    alert("You can't add an empty task");
  }
  console.log(tasks);
};

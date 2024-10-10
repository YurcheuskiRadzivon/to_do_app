document.addEventListener("DOMContentLoaded", function () {
  tasks.forEach((task) => addTask(task.name, task.description));
  
  listOfTasks.addEventListener('click', function(event) {
    const target = event.target.closest('button'); // Найти ближайшую кнопку

    if (target && target.classList.contains('action_button')) {
      const actionsPanel = target.nextElementSibling;
      actionsPanel.style.display = actionsPanel.style.display === 'none' ? 'block' : 'none';
      target.style.display = 'none'; // Скрыть кнопку действий
    }

    if (target && target.id === 'delete') {
      const taskElement = target.closest('.task_line');
      const taskIndex = taskElement.querySelector('.task').dataset.ind;
      tasks.splice(taskIndex, 1);
      saveTasks();
      taskElement.remove();
    }

    if (target && target.id === 'done') {
      const taskElement = target.closest('.task_line');
      taskElement.querySelector('.content').style.textDecoration = 'line-through';
    }

    if (target && target.id === 'rename') {
      const taskElement = target.closest('.task_line');
      const newName = prompt("Enter new task name:");
      const newDesc = prompt("Enter new task description:");
      if (newName && newDesc) {
        const taskIndex = taskElement.querySelector('.task').dataset.ind;
        tasks[taskIndex].name = newName;
        tasks[taskIndex].description = newDesc;
        saveTasks();
        taskElement.querySelector('.heading').textContent = newName;
        taskElement.querySelector('.para').textContent = newDesc;
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
                <div id="line" class="task" data-ind="${ind++}">
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
                  <button id="done" class="action_with_task">
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      x="0px"
                      y="0px"
                      width="40"
                      height="40"
                      viewBox="0 0 24 24"
                    >
                      <path
                        fill="#fafafa" d="M 20.292969 5.2929688 L 9 16.585938 L 4.7070312 12.292969 L 3.2929688 13.707031 L 9 19.414062 L 21.707031 6.7070312 L 20.292969 5.2929688 z"
                      ></path>
                    </svg>
                  </button>

                  <button id="delete" class="action_with_task">
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      x="0px"
                      y="0px"
                      width="40"
                      height="40"
                      viewBox="0 0 50 50"
                    >
                      <path
                        fill="#fafafa" d="M 21 0 C 19.355469 0 18 1.355469 18 3 L 18 5 L 10.1875 5 C 10.0625 4.976563 9.9375 4.976563 9.8125 5 L 8 5 C 7.96875 5 7.9375 5 7.90625 5 C 7.355469 5.027344 6.925781 5.496094 6.953125 6.046875 C 6.980469 6.597656 7.449219 7.027344 8 7 L 9.09375 7 L 12.6875 47.5 C 12.8125 48.898438 14.003906 50 15.40625 50 L 34.59375 50 C 35.996094 50 37.1875 48.898438 37.3125 47.5 L 40.90625 7 L 42 7 C 42.359375 7.003906 42.695313 6.816406 42.878906 6.503906 C 43.058594 6.191406 43.058594 5.808594 42.878906 5.496094 C 42.695313 5.183594 42.359375 4.996094 42 5 L 32 5 L 32 3 C 32 1.355469 30.644531 0 29 0 Z M 21 2 L 29 2 C 29.5625 2 30 2.4375 30 3 L 30 5 L 20 5 L 20 3 C 20 2.4375 20.4375 2 21 2 Z M 11.09375 7 L 38.90625 7 L 35.3125 47.34375 C 35.28125 47.691406 34.910156 48 34.59375 48 L 15.40625 48 C 15.089844 48 14.71875 47.691406 14.6875 47.34375 Z M 18.90625 9.96875 C 18.863281 9.976563 18.820313 9.988281 18.78125 10 C 18.316406 10.105469 17.988281 10.523438 18 11 L 18 44 C 17.996094 44.359375 18.183594 44.695313 18.496094 44.878906 C 18.808594 45.058594 19.191406 45.058594 19.503906 44.878906 C 19.816406 44.695313 20.003906 44.359375 20 44 L 20 11 C 20.011719 10.710938 19.894531 10.433594 19.6875 10.238281 C 19.476563 10.039063 19.191406 9.941406 18.90625 9.96875 Z M 24.90625 9.96875 C 24.863281 9.976563 24.820313 9.988281 24.78125 10 C 24.316406 10.105469 23.988281 10.523438 24 11 L 24 44 C 23.996094 44.359375 24.183594 44.695313 24.496094 44.878906 C 24.808594 45.058594 25.191406 45.058594 25.503906 44.878906 C 25.816406 44.695313 26.003906 44.359375 26 44 L 26 11 C 26.011719 10.710938 25.894531 10.433594 25.6875 10.238281 C 25.476563 10.039063 25.191406 9.941406 24.90625 9.96875 Z M 30.90625 9.96875 C 30.863281 9.976563 30.820313 9.988281 30.78125 10 C 30.316406 10.105469 29.988281 10.523438 30 11 L 30 44 C 29.996094 44.359375 30.183594 44.695313 30.496094 44.878906 C 30.808594 45.058594 31.191406 45.058594 31.503906 44.878906 C 31.816406 44.695313 32.003906 44.359375 32 44 L 32 11 C 32.011719 10.710938 31.894531 10.433594 31.6875 10.238281 C 31.476563 10.039063 31.191406 9.941406 30.90625 9.96875 Z"
                      ></path>
                    </svg>
                  </button>

                  <button id="rename" class="action_with_task">
                    <svg fill="#000000" height="40px" width="40px" version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" 
                    viewBox="0 0 512 512" xml:space="preserve">
                 <g>
                   <g>
                     <path fill="#fafafa" d="M104.791,392.054c-0.148-0.727-0.362-1.442-0.646-2.123c-0.284-0.693-0.637-1.351-1.045-1.964
                       c-0.409-0.625-0.886-1.204-1.408-1.726s-1.101-0.999-1.727-1.408c-0.625-0.409-1.283-0.761-1.963-1.045s-1.397-0.5-2.123-0.647
                       c-1.465-0.295-2.975-0.295-4.44,0c-0.727,0.148-1.442,0.363-2.123,0.647s-1.34,0.636-1.963,1.045
                       c-0.613,0.409-1.192,0.886-1.727,1.408c-0.522,0.522-0.999,1.101-1.407,1.726c-0.409,0.613-0.761,1.272-1.045,1.964
                       c-0.284,0.681-0.5,1.397-0.647,2.123s-0.216,1.476-0.216,2.214c-0.002,2.998,1.203,5.916,3.313,8.028
                       c0.535,0.522,1.113,0.999,1.727,1.408c0.625,0.42,1.283,0.761,1.963,1.056c0.682,0.284,1.397,0.5,2.123,0.647
                       c0.738,0.148,1.476,0.216,2.226,0.216c2.986,0,5.916-1.215,8.028-3.327c2.112-2.112,3.327-5.03,3.327-8.028
                       C105.018,393.531,104.94,392.781,104.791,392.054z"/>
                   </g>
                 </g>
                 <g>
                   <g>
                     <path fill="#fafafa" d="M475.367,131.818l16.938-16.938c26.238-26.239,26.238-68.931,0-95.17C479.595,7,462.695,0,444.72,0
                       s-34.874,7-47.584,19.71l-16.939,16.939c-2.127-2.109-4.998-3.294-7.995-3.294c-3.011,0-5.899,1.197-8.029,3.326L93.884,306.971
                       c-0.002,0.002-0.003,0.003-0.006,0.006l-2.333,2.333c-0.595,0.596-1.122,1.256-1.573,1.968l-58.996,93.483
                       c-2.831,4.487-2.178,10.338,1.574,14.089l0.083,0.083l-31.767,77.4c-1.729,4.214-0.774,9.054,2.429,12.294
                       C5.471,510.827,8.395,512,11.374,512c1.408,0,2.828-0.262,4.188-0.802l78.316-31.089c2.09,1.766,4.693,2.682,7.32,2.682
                       c2.089,0,4.19-0.575,6.057-1.752l93.484-58.996c0.713-0.45,1.373-0.978,1.969-1.574l2.335-2.335l262.263-262.263l3.444,3.444
                       c3.536,3.536,5.483,8.236,5.483,13.236c0,5.001-1.947,9.701-5.483,13.236L353.106,303.435c-4.434,4.434-4.434,11.624,0,16.058
                       c2.218,2.218,5.123,3.326,8.029,3.326s5.811-1.109,8.029-3.326L486.806,201.85c7.826-7.825,12.135-18.229,12.135-29.296
                       s-4.309-21.471-12.135-29.295L475.367,131.818z M32.025,480.23l18.013-43.891l26.3,26.3L32.025,480.23z M102.809,456.991
                       l-47.784-47.784l48.374-76.654l17.958,17.958l-12.205,12.205c-4.434,4.434-4.434,11.624,0,16.059
                       c2.218,2.218,5.123,3.326,8.029,3.326c2.906,0,5.811-1.108,8.029-3.326l12.205-12.205l42.046,42.046L102.809,456.991z
                        M197.013,394.048l-43.537-43.537l205.947-205.947c4.434-4.434,4.434-11.624,0-16.059c-4.435-4.434-11.623-4.434-16.059,0
                       L137.417,334.453l-19.449-19.449L372.202,60.769l72.753,72.753l6.292,6.292L197.013,394.048z M459.31,115.76l-63.053-63.053
                       l16.937-16.939c8.421-8.421,19.616-13.058,31.526-13.058s23.106,4.637,31.527,13.059c17.384,17.383,17.384,45.67,0,63.053
                       L459.31,115.76z"/>
                   </g>
                 </g>
                 </svg>
                  </button>
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

const taskManagerContainer = document.querySelector('.taskManager');
const confirmEl = document.querySelector('.confirm');
const confirmedBtn = confirmEl.querySelector('.confirmed');
const cancelledBtn = confirmEl.querySelector('.cancel');
let indexToBeDeleted = null;

const Tasks = [
  {
    id: String,
    name: String,
    description: String,
    isCompleted: Boolean,
    createdAt: Date,
    updatedAt: Date,
    author: String,
    comments: Array,
    shared: Boolean,
  },
];

let allTasks = [...Tasks];

const form = document.getElementById('taskForm');
const submitButton = document.getElementById('submit');

// Add event listener to the form submit event
form.addEventListener('submit', handleFormSubmit);

// Add event listener to the button click event
submitButton.addEventListener('click', handleFormSubmit);

// Function to handle form submission
async function handleFormSubmit(event) {
  event.preventDefault();
  const authorNameInput = document.getElementById('authorName');
  const taskNameInput = document.getElementById('taskName');
  const taskDescriptionInput = document.getElementById('taskDescription');

  const authorName = authorNameInput.value.trim();
  const taskName = taskNameInput.value.trim();
  const taskDescription = taskDescriptionInput.value.trim();

  if (authorName !== '' && taskName !== '') {
    // Create a new task object
    const newTask = {
      name: taskName,
      description: taskDescription,
      isCompleted: false,
      createdAt: new Date().toISOString(),
      updatedAt: null,
      author: authorName,
      comments: [],
      shared: false,
    };

    try {
      // Make API call to create task
      await createTask(newTask);

      // Render tasks after creating a new task
      renderTasks();
    } catch (error) {
      console.error('Error creating task:', error);
      // Handle error as needed
    }
  }
  // Clear input fields
  authorNameInput.value = '';
  taskNameInput.value = '';
  taskDescriptionInput.value = '';
}

// Function to handle button click (optional)
function handleButtonClick(event) {
  // You can add additional logic here if needed
  console.log('Button clicked!');
}

// Function to make API call and create a new task
async function createTask(task) {
  try {
    const response = await fetch(
      'https://asim-task-manage.vercel.app/api/create',
      {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(task),
      },
    );

    if (!response.ok) {
      throw new Error('Error creating task');
    }
  } catch (error) {
    console.error(error.message);
  }
}

// Initial rendering of tasks
renderTasks();

// Function to render tasks
async function renderTasks() {
  const taskContainer = document.getElementById('taskContainer');
  taskContainer.innerHTML = '';

  try {
    // Make API call to get tasks
    const response = await fetch('https://asim-task-manage.vercel.app/api/get');
    if (!response.ok) {
      throw new Error('Error fetching tasks Yeah');
    }

    allTasks = await response.json();

    allTasks.data.forEach((task, index) => {
      const taskCard = document.createElement('div');
      taskCard.classList.add('taskCard');
      let classVal = 'pending';
      let textVal = 'Pending';
      if (task.isCompleted) {
        classVal = 'completed';
        textVal = 'Completed';
      }
      taskCard.classList.add(classVal);

      const taskText = document.createElement('p');
      taskText.innerText = task.name;

      const taskStatus = document.createElement('p');
      taskStatus.classList.add('status');
      taskStatus.innerText = textVal;

      const toggleButton = document.createElement('button');
      toggleButton.classList.add('button-box');
      const btnContentEl = document.createElement('span');
      btnContentEl.classList.add('green');
      btnContentEl.innerText = task.isCompleted
        ? 'Mark as Pending'
        : 'Mark as Completed';
      toggleButton.appendChild(btnContentEl);
      toggleButton.addEventListener('click', async () => {
        task.updatedAt = new Date().toISOString();
        task.isCompleted = !task.isCompleted;
        await updateTask(task);
        renderTasks();
      });

      const deleteButton = document.createElement('button');
      deleteButton.classList.add('button-box');
      const delBtnContentEl = document.createElement('span');
      delBtnContentEl.classList.add('red');
      delBtnContentEl.innerText = 'Delete';
      deleteButton.appendChild(delBtnContentEl);
      deleteButton.addEventListener('click', () => {
        indexToBeDeleted = task.id;
        confirmEl.style.display = 'block';
        taskManagerContainer.classList.add('overlay');
      });

      const commentsSection = document.createElement('div');
      commentsSection.classList.add('comments-section');
      const commentsIcon = document.createElement('div');
      commentsIcon.classList.add('comments-icon');
      commentsIcon.innerHTML = '&#128172;'; // Comment icon
      commentsIcon.addEventListener('click', () =>
        toggleComments(commentsContainer),
      );

      const commentsContainer = document.createElement('div');
      commentsContainer.classList.add('comments-container');
      //Add a heading Comments
      const commentsHeading = document.createElement('h4');
      commentsHeading.innerText = 'Comments';
      commentsContainer.appendChild(commentsHeading);

      // Display comments
      task.comments.forEach((comment) => {
        const commentElement = document.createElement('p');
        commentElement.innerText = comment;
        commentsContainer.appendChild(commentElement);
      });

      // Add input for new comment
      const commentInput = document.createElement('input');
      commentInput.placeholder = 'Add a comment...';
      commentInput.classList.add('task-input'); // Apply the same styling as the form inputs

      // Add button to submit comment
      const addCommentButton = document.createElement('button');
      addCommentButton.innerText = 'Add Comment';
      addCommentButton.classList.add('task-submit-button'); // Apply the same styling as the form submit button
      addCommentButton.addEventListener('click', async () => {
        const newComment = commentInput.value.trim();
        if (newComment !== '') {
          task.comments.push(newComment);
          await updateTask(task);
          renderTasks();
        }
      });

      commentsContainer.appendChild(commentInput);
      commentsContainer.appendChild(addCommentButton);

      taskCard.appendChild(commentsIcon);
      taskCard.appendChild(commentsContainer);

      taskCard.appendChild(taskText);
      taskCard.appendChild(taskStatus);
      taskCard.appendChild(commentsContainer);
      taskCard.appendChild(toggleButton);
      taskCard.appendChild(deleteButton);

      taskContainer.appendChild(taskCard);
    });
  } catch (error) {
    console.error(error.message);
  }
}

function toggleComments(commentsContainer) {
  commentsContainer.classList.toggle('comments-visible');
}

// Function to make API call and update task status
async function updateTask(task) {
  try {
    const response = await fetch(
      `https://asim-task-manage.vercel.app/api/update/${task.id}`,
      {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(task),
      },
    );

    if (!response.ok) {
      throw new Error('Error updating task status');
    }
  } catch (error) {
    console.error(error.message);
  }
}

// Function to delete the selected task
async function deleteTask(index) {
  try {
    const response = await fetch(
      `https://asim-task-manage.vercel.app/api/delete/${index}`,
      {
        method: 'DELETE',
      },
    );

    if (!response.ok) {
      throw new Error('Error deleting task');
    }

    renderTasks();
  } catch (error) {
    console.error(error.message);
  }
}

confirmedBtn.addEventListener('click', async () => {
  confirmEl.style.display = 'none';
  taskManagerContainer.classList.remove('overlay');
  await deleteTask(indexToBeDeleted);
});

cancelledBtn.addEventListener('click', () => {
  confirmEl.style.display = 'none';
  taskManagerContainer.classList.remove('overlay');
});

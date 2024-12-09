document.addEventListener('DOMContentLoaded', function () {
    const menuBtn = document.querySelector('.menu-btn');
    const sidebar = document.querySelector('.sidebar');
    const editBtn = document.querySelector('.edit-btn');
    const saveBtn = document.querySelector('.save-btn');
    const cancelBtn = document.querySelector('.cancel-btn');
    const deleteBtn = document.querySelector('.delete-btn');
    const infoValues = document.querySelectorAll('.info-value');
    const infoInputs = document.querySelectorAll('.info-input');

    // Store original values
    const originalValues = Array.from(infoValues).map(value => value.textContent.trim());

    // Toggle sidebar
    menuBtn.addEventListener('click', function () {
        sidebar.classList.toggle('active');
    });

    // Close sidebar when clicking outside
    document.addEventListener('click', function (e) {
        if (!sidebar.contains(e.target) && !menuBtn.contains(e.target)) {
            sidebar.classList.remove('active');
        }
    });

    // Edit button functionality
    editBtn.addEventListener('click', function () {
        infoValues.forEach((value, index) => {
            if (index < 3) { // Skip the Added Time field
                value.classList.add('hidden');
            }
        });
        infoInputs.forEach(input => input.classList.add('active'));
        saveBtn.classList.add('active');
        cancelBtn.classList.add('active');
        editBtn.style.display = 'none';
    });

    // Save button functionality
    saveBtn.addEventListener('click', async function () {
        try {
            const taskElement = document.querySelector('.task-id');
            const taskId = taskElement.getAttribute('data-id');
            const taskData = {
                title: infoInputs[0].value,
                description: infoInputs[1].value,
                status: infoInputs[2].value === 'true'
            };


            const response = await fetch(`/tasks/${taskId}`, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(taskData)
            });

            if (!response.ok) {
                throw new Error('Failed to update task');
            }

            // Update the displayed values
            infoValues[0].textContent = taskData.title;
            infoValues[1].textContent = taskData.description;
            const statusBadge = infoValues[2].querySelector('.status-badge');
            statusBadge.textContent = taskData.status.toString();
            statusBadge.className = `status-badge status-${taskData.status}`;

            resetEditingState();
            alert('Task updated successfully!');

            window.location.href = '/tasks';
        } catch (error) {
            console.error('Error:', error);
            alert('Error updating task');
        }
    });

    // Cancel button functionality
    cancelBtn.addEventListener('click', function () {
        resetEditingState();
        infoInputs[0].value = originalValues[0];
        infoInputs[1].value = originalValues[1];
        infoInputs[2].value = 'false'; // Reset to default status
    });


    deleteBtn.addEventListener('click', async function () {
        if (confirm('Are you sure you want to delete this task?')) {
            try {
                const taskElement = document.querySelector('.task-id');
                const taskId = taskElement.getAttribute('data-id');
                const response = await fetch(`/tasks/${taskId}`, {
                    method: 'DELETE',
                    headers: {
                        'Content-Type': 'application/json',
                    }
                });

                if (!response.ok) {
                    throw new Error('Failed to delete task');
                }

                alert('Task deleted successfully!');

                window.location.href = '/tasks';
            } catch (error) {
                console.error('Error:', error);
                alert('Error deleting task');
            }
        }
    });

    function resetEditingState() {
        infoValues.forEach((value, index) => {
            if (index < 3) { // Skip the Added Time field
                value.classList.remove('hidden');
            }
        });
        infoInputs.forEach(input => input.classList.remove('active'));
        saveBtn.classList.remove('active');
        cancelBtn.classList.remove('active');
        editBtn.style.display = 'block';
    }


});
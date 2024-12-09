document.addEventListener('DOMContentLoaded', function () {
    const menuBtn = document.querySelector('.menu-btn');
    const sidebar = document.querySelector('.sidebar');
    const imageUpload = document.getElementById('imageUpload');
    const taskForm = document.getElementById('taskForm');
    let uploadedImages = [];

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


    function createImagePreview() {
        if (document.querySelectorAll('.image-preview').length >= 3) return;

        const preview = document.createElement('div');
        preview.className = 'image-preview';

        const input = document.createElement('input');
        input.type = 'file';
        input.accept = 'image/*';

        const span = document.createElement('span');
        span.textContent = '+';

        preview.appendChild(input);
        preview.appendChild(span);

        preview.addEventListener('click', () => input.click());

        input.addEventListener('change', function (e) {
            if (this.files && this.files[0]) {
                const reader = new FileReader();
                reader.onload = function (e) {
                    const img = document.createElement('img');
                    img.src = e.target.result;
                    img.className = 'preview-image';

                    const removeBtn = document.createElement('button');
                    removeBtn.className = 'remove-image';
                    removeBtn.textContent = '×';
                    removeBtn.onclick = function (e) {
                        e.stopPropagation();
                        preview.remove();
                        uploadedImages = uploadedImages.filter(image => image !== img.src);
                        createImagePreview();
                    };

                    preview.innerHTML = '';
                    preview.appendChild(img);
                    preview.appendChild(removeBtn);

                    uploadedImages.push(e.target.result);
                    createImagePreview();
                };
                reader.readAsDataURL(this.files[0]);
            }
        });

        imageUpload.appendChild(preview);
    }

    createImagePreview();


    taskForm.addEventListener('submit', async function (e) {
        e.preventDefault();

        const formData = {
            title: taskForm.querySelector('[name="title"]').value,
            description: taskForm.querySelector('[name="description"]').value,
            status: taskForm.querySelector('[name="status"]').value,

        };

        try {
            const response = await fetch('/tasks', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(formData)
            });

            if (response.ok) {
                window.location.href = '/tasks';
            } else {
                throw new Error('Failed to create task');
            }
        } catch (error) {
            console.error('Error:', error);
            alert('Error creating task');
        }
    });
});
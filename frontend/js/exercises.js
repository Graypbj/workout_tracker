// frontend/js/exercises.js
import { apiUrlBase, getToken, handleUnauthorized, fetchWithAuth } from './api.js';

document.addEventListener('DOMContentLoaded', () => {
    const exerciseForm = document.getElementById('manage-exercise-form');
    const exerciseIdInput = document.getElementById('exercise-id');
    const exerciseNameInput = document.getElementById('exercise-name');
    const exerciseDescriptionInput = document.getElementById('exercise-description');
    const exerciseCategoryInput = document.getElementById('exercise-category');
    const exercisesListElement = document.getElementById('exercises-list');
    const cancelEditBtn = document.getElementById('cancel-edit-btn');
    const errorMessageElement = document.getElementById('exercise-error-message');
    const logoutLink = document.getElementById('logout-link');

    const displayErrorMessage = (message) => {
        if(errorMessageElement) {
            errorMessageElement.textContent = message;
            errorMessageElement.style.display = 'block';
        } else {
            console.error("Error message element not found for:", message);
        }
    };

    const clearErrorMessage = () => {
        if(errorMessageElement) {
            errorMessageElement.textContent = '';
            errorMessageElement.style.display = 'none';
        }
    };

    const resetForm = () => {
        if(exerciseForm) exerciseForm.reset();
        if(exerciseIdInput) exerciseIdInput.value = '';
        if(cancelEditBtn) cancelEditBtn.style.display = 'none';
        clearErrorMessage();
    };

    const displayExercises = async () => {
        clearErrorMessage();
        if (!getToken()) {
            handleUnauthorized();
            return;
        }
        try {
            const response = await fetchWithAuth(`${apiUrlBase}/exercises`);
            // For DELETE requests, a 204 No Content is a success but has no body to parse.
            if (response.status === 204) {
                 if(exercisesListElement) exercisesListElement.innerHTML = '<li>No exercises found or exercises deleted.</li>';
                return;
            }
            const data = await response.json();
            const exercises = data.exercises || [];

            if (!exercisesListElement) return;
            exercisesListElement.innerHTML = '';

            if (exercises.length === 0) {
                exercisesListElement.innerHTML = '<li>No exercises found. Add some!</li>';
                return;
            }

            exercises.forEach(exercise => {
                const listItem = document.createElement('li');
                listItem.innerHTML = `
                    <strong>${exercise.name}</strong>
                    <p>${exercise.description || 'No description'}</p>
                    <small>Category: ${exercise.category || 'N/A'}</small>
                    <div>
                        <button class="edit-btn" data-id="${exercise.id}">Edit</button>
                        <button class="delete-btn" data-id="${exercise.id}">Delete</button>
                    </div>
                `;
                exercisesListElement.appendChild(listItem);

                listItem.querySelector('.edit-btn').addEventListener('click', () => populateFormForEdit(exercise));
                listItem.querySelector('.delete-btn').addEventListener('click', () => deleteExercise(exercise.id));
            });
        } catch (error) {
            console.error('Failed to display exercises:', error);
            if (error.message !== 'Unauthorized' && error.message !== 'No token found, authorization required.') {
                displayErrorMessage(`Failed to load exercises: ${error.message}`);
            }
        }
    };

    const populateFormForEdit = (exercise) => {
        if(exerciseIdInput) exerciseIdInput.value = exercise.id;
        if(exerciseNameInput) exerciseNameInput.value = exercise.name;
        if(exerciseDescriptionInput) exerciseDescriptionInput.value = exercise.description || '';
        if(exerciseCategoryInput) exerciseCategoryInput.value = exercise.category || '';
        if(cancelEditBtn) cancelEditBtn.style.display = 'inline-block';
        if(exerciseNameInput) exerciseNameInput.focus();
        clearErrorMessage();
    };

    if (exerciseForm) {
        exerciseForm.addEventListener('submit', async (e) => {
            e.preventDefault();
            clearErrorMessage();

            const id = exerciseIdInput ? exerciseIdInput.value : null;
            const name = exerciseNameInput ? exerciseNameInput.value.trim() : '';
            const description = exerciseDescriptionInput ? exerciseDescriptionInput.value.trim() : '';
            const category = exerciseCategoryInput ? exerciseCategoryInput.value.trim() : '';

            if (!name) {
                displayErrorMessage('Exercise name is required.');
                return;
            }

            const exerciseData = { name, description, category };
            const method = id ? 'PUT' : 'POST';
            if (id) {
                exerciseData.id = id;
            }

            try {
                await fetchWithAuth(`${apiUrlBase}/exercises`, {
                    method: method,
                    body: JSON.stringify(exerciseData)
                });
                resetForm();
                await displayExercises();
            } catch (error) {
                console.error('Failed to save exercise:', error);
                if (error.message !== 'Unauthorized' && error.message !== 'No token found, authorization required.') {
                    displayErrorMessage(`Failed to save exercise: ${error.message}`);
                }
            }
        });
    }

    const deleteExercise = async (id) => {
        if (!confirm('Are you sure you want to delete this exercise?')) {
            return;
        }
        clearErrorMessage();
        try {
            await fetchWithAuth(`${apiUrlBase}/exercises`, {
                method: 'DELETE',
                body: JSON.stringify({ id: id })
            });
            await displayExercises();
        } catch (error) {
            console.error('Failed to delete exercise:', error);
            if (error.message !== 'Unauthorized' && error.message !== 'No token found, authorization required.') {
                displayErrorMessage(`Failed to delete exercise: ${error.message}`);
            }
        }
    };

    if(cancelEditBtn) {
        cancelEditBtn.addEventListener('click', () => {
            resetForm();
        });
    }

    if (logoutLink) {
        logoutLink.addEventListener('click', (e) => {
            e.preventDefault();
            handleUnauthorized();
        });
    }

    if (window.location.pathname.endsWith('exercises.html')) {
         if (getToken()) {
            displayExercises();
        } else {
            handleUnauthorized();
        }
    }
});

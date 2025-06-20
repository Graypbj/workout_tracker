// frontend/js/workouts.js
import { apiUrlBase, getToken, handleUnauthorized, fetchWithAuth } from './api.js';

document.addEventListener('DOMContentLoaded', () => {
    const workoutsListElement = document.getElementById('workouts-list');
    const createWorkoutForm = document.getElementById('create-workout-form');
    const workoutTypeInput = document.getElementById('workout-type');
    const workoutNotesInput = document.getElementById('workout-notes');
    const logoutLink = document.getElementById('logout-link');
    const errorMessageElement = document.createElement('p'); // For displaying errors
    errorMessageElement.style.color = 'red';
    if (createWorkoutForm) {
        createWorkoutForm.parentNode.insertBefore(errorMessageElement, createWorkoutForm.nextSibling);
    }


    const displayErrorMessage = (message) => {
        errorMessageElement.textContent = message;
    };

    const clearErrorMessage = () => {
        errorMessageElement.textContent = '';
    };

    const displayWorkouts = async () => {
        clearErrorMessage();
        if (!getToken()) { // Check token before attempting to fetch
            handleUnauthorized();
            return;
        }
        try {
            const response = await fetchWithAuth(`${apiUrlBase}/workouts`);
            // For DELETE requests, a 204 No Content is a success but has no body to parse.
            if (response.status === 204) {
                workoutsListElement.innerHTML = '<li>No workouts found or workouts deleted.</li>'; // Or handle as appropriate
                return;
            }
            const data = await response.json();
            const workouts = data.workouts || [];

            workoutsListElement.innerHTML = ''; // Clear existing list

            if (workouts.length === 0) {
                workoutsListElement.innerHTML = '<li>No workouts found. Add some!</li>';
                return;
            }

            workouts.forEach(workout => {
                const listItem = document.createElement('li');
                const notes = workout.notes || 'No notes';
                const date = new Date(workout.workout_date).toLocaleDateString();
                listItem.innerHTML = `
                    <strong>${workout.workout_type}</strong> - ${date}
                    <p>${notes}</p>
                    <div>
                        <button class="delete-btn" data-id="${workout.id}">Delete</button>
                    </div>
                `;
                workoutsListElement.appendChild(listItem);

                listItem.querySelector('.delete-btn').addEventListener('click', () => deleteWorkout(workout.id));
            });
        } catch (error) {
            console.error('Failed to display workouts:', error);
            if (error.message !== 'Unauthorized' && error.message !== 'No token found, authorization required.') {
                 displayErrorMessage(`Error loading workouts: ${error.message}`);
            }
        }
    };

    const deleteWorkout = async (workoutId) => {
        clearErrorMessage();
        if (!confirm('Are you sure you want to delete this workout?')) return;

        try {
            await fetchWithAuth(`${apiUrlBase}/workouts`, {
                method: 'DELETE',
                body: JSON.stringify({ id: workoutId })
            });
            await displayWorkouts(); // Refresh list
        } catch (error) {
            console.error('Failed to delete workout:', error);
             if (error.message !== 'Unauthorized' && error.message !== 'No token found, authorization required.') {
                displayErrorMessage(`Error deleting workout: ${error.message}`);
            }
        }
    };

    if (createWorkoutForm) {
        createWorkoutForm.addEventListener('submit', async (e) => {
            e.preventDefault();
            clearErrorMessage();
            const type = workoutTypeInput.value.trim();
            const notes = workoutNotesInput.value.trim();

            if (!type) {
                displayErrorMessage('Workout type is required.');
                return;
            }

            try {
                await fetchWithAuth(`${apiUrlBase}/workouts`, {
                    method: 'POST',
                    body: JSON.stringify({ workout_type: type, notes: notes })
                });
                createWorkoutForm.reset();
                await displayWorkouts(); // Refresh list
            } catch (error) {
                console.error('Failed to create workout:', error);
                if (error.message !== 'Unauthorized' && error.message !== 'No token found, authorization required.') {
                    displayErrorMessage(`Error creating workout: ${error.message}`);
                }
            }
        });
    }

    if (logoutLink) {
        logoutLink.addEventListener('click', (e) => {
            e.preventDefault();
            handleUnauthorized();
        });
    }

    // Initial load
    if (window.location.pathname.endsWith('workouts.html')) { // Ensure this runs only on workouts.html
         if (getToken()) {
            displayWorkouts();
        } else {
            handleUnauthorized();
        }
    }
});

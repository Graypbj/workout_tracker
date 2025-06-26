// workouts.js
const workoutsList = document.getElementById("workouts-list");
const form = document.getElementById("create-workout-form");
const API_URL = "http://localhost:8080/api";

async function fetchWorkouts() {
	// Placeholder for future GET request implementation
}

form.addEventListener("submit", async (e) => {
	e.preventDefault();

	const type = document.getElementById("type").value;
	const notes = document.getElementById("workout-notes").value;

	const res = await fetch(`${API_URL}/workouts`, {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
			Authorization: `Bearer ${localStorage.getItem("token")}`,
		},
		body: JSON.stringify({
			workout_type: type,
			notes,
		}),
	});

	if (!res.ok) {
		alert("Error creating workout");
		return;
	}

	const workout = await res.json();
	form.reset();
	renderWorkout(workout);
});

function renderWorkout(workout) {
	const container = document.createElement("div");
	container.classList.add("workout");

	container.innerHTML = `
    <h4>${workout.workout_type.replace("_", " ")}</h4>
    <p>${workout.notes || "No notes"}</p>
    <p><small>${new Date(workout.created_at).toLocaleDateString()}</small></p>
    <div class="exercises"></div>
  `;

	const formTemplate = document.getElementById("exercise-form-template");
	const formClone = formTemplate.content.cloneNode(true);
	const exerciseForm = formClone.querySelector("form");

	const strengthInputs = exerciseForm.querySelector(".strength-inputs");
	const cardioInputs = exerciseForm.querySelector(".cardio-inputs");

	const exerciseTypeSelect = exerciseForm.querySelector("select[name='exercise_type']");
	exerciseTypeSelect.addEventListener("change", (e) => {
		if (e.target.value === "strength_training") {
			strengthInputs.classList.remove("hidden");
			cardioInputs.classList.add("hidden");
		} else if (e.target.value === "cardio") {
			cardioInputs.classList.remove("hidden");
			strengthInputs.classList.add("hidden");
		} else {
			strengthInputs.classList.add("hidden");
			cardioInputs.classList.add("hidden");
		}
	});

	exerciseForm.addEventListener("submit", async (e) => {
		e.preventDefault();
		const formData = new FormData(exerciseForm);
		const name = formData.get("name");
		const type = formData.get("exercise_type");

		const exerciseRes = await fetch(`${API_URL}/exercises`, {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
				Authorization: `Bearer ${localStorage.getItem("token")}`,
			},
			body: JSON.stringify({ name, exercise_type: type }),
		});

		if (!exerciseRes.ok) {
			alert("Error creating exercise");
			return;
		}

		const exercise = await exerciseRes.json();

		if (type === "strength_training") {
			const setNumber = parseInt(formData.get("set_number"));
			const reps = parseInt(formData.get("reps"));
			const weight = formData.get("weight");

			const sessionRes = await fetch(`${API_URL}/strength_training_sessions`, {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
					Authorization: `Bearer ${localStorage.getItem("token")}`,
				},
				body: JSON.stringify({
					workout_id: workout.id,
					exercise_id: exercise.id,
					notes: "",
				}),
			});

			const sessionData = await sessionRes.json();
			const sessionID = sessionData.strength_training_session.id;

			await fetch(`${API_URL}/strength_training_sets`, {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
					Authorization: `Bearer ${localStorage.getItem("token")}`,
				},
				body: JSON.stringify({
					session_id: sessionID,
					set_number: setNumber,
					reps,
					weight,
				}),
			});
		} else if (type === "cardio") {
			const distance = parseFloat(formData.get("distance"));
			const duration = parseInt(formData.get("duration"));
			const unit = formData.get("distance_unit");

			// Convert to duration string format: e.g., "30m" or "45m"
			const timeStr = `${duration}m`;

			await fetch(`${API_URL}/cardio_training_sessions`, {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
					Authorization: `Bearer ${localStorage.getItem("token")}`,
				},
				body: JSON.stringify({
					workout_id: workout.id,
					exercise_id: exercise.id,
					distance,
					time: timeStr,
					notes: `${distance} ${unit} in ${timeStr}`,
				}),
			});
		}

		const exercisesContainer = container.querySelector(".exercises");
		const p = document.createElement("p");
		p.textContent = `✅ ${exercise.name} (${type})`;
		exercisesContainer.appendChild(p);

		exerciseForm.reset();
		strengthInputs.classList.add("hidden");
		cardioInputs.classList.add("hidden");
	});

	container.appendChild(formClone);
	workoutsList.prepend(container);
}

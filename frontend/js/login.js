document.addEventListener('DOMContentLoaded', () => {
	const form = document.getElementById('login-form');

	form.addEventListener('submit', async (e) => {
		e.preventDefault();

		const email = document.getElementById('email').value.trim();
		const password = document.getElementById('password').value;

		try {
			const response = await fetch('http://localhost:8080/api/login', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({ email, password })
			});

			if (response.ok) {
				const data = await response.json();

				localStorage.setItem('token', data.token);
				localStorage.setItem('refreshToken', data.refresh_token);

				window.location.href = '/workouts.html';
			}
		} catch (err) {
			console.error('Login request failed:', err);
		}
	});
});

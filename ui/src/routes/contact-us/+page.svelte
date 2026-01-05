<script lang="ts">
	let name = $state('');
	let email = $state('');
	let subject = $state('');
	let message = $state('');
	let isSubmitting = $state(false);
	let submitStatus = $state<{ type: 'success' | 'error'; message: string } | null>(null);

	async function handleSubmit(e: Event) {
		e.preventDefault();
		isSubmitting = true;
		submitStatus = null;

		try {
			const response = await fetch('/api/contact', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({ name, email, subject, message })
			});

			const data = await response.json();

			if (data.success) {
				submitStatus = { type: 'success', message: data.message };
				// Clear form
				name = '';
				email = '';
				subject = '';
				message = '';
			} else {
				submitStatus = { type: 'error', message: data.message };
			}
		} catch (error) {
			submitStatus = { 
				type: 'error', 
				message: 'Failed to send message. Please try again later.' 
			};
		} finally {
			isSubmitting = false;
		}
	}
</script>

<div class="contact-container">
	<section class="hero">
		<h1>Contact Us</h1>
		<p class="subtitle">We'd love to hear from you. Send us a message!</p>
	</section>

	<section class="form-section">
		{#if submitStatus}
			<div class="alert {submitStatus.type === 'success' ? 'alert-success' : 'alert-error'}">
				{submitStatus.message}
			</div>
		{/if}

		<form onsubmit={handleSubmit}>
			<div class="form-group">
				<label for="name">Name *</label>
				<input
					type="text"
					id="name"
					bind:value={name}
					required
					placeholder="Your name"
					disabled={isSubmitting}
				/>
			</div>

			<div class="form-group">
				<label for="email">Email *</label>
				<input
					type="email"
					id="email"
					bind:value={email}
					required
					placeholder="your.email@example.com"
					disabled={isSubmitting}
				/>
			</div>

			<div class="form-group">
				<label for="subject">Subject</label>
				<input
					type="text"
					id="subject"
					bind:value={subject}
					placeholder="What is this about?"
					disabled={isSubmitting}
				/>
			</div>

			<div class="form-group">
				<label for="message">Message *</label>
				<textarea
					id="message"
					bind:value={message}
					required
					rows="6"
					placeholder="Tell us what's on your mind..."
					disabled={isSubmitting}
				></textarea>
			</div>

			<button type="submit" disabled={isSubmitting} class="submit-btn">
				{isSubmitting ? 'Sending...' : 'Send Message'}
			</button>
		</form>
	</section>
</div>

<style>
	.contact-container {
		max-width: 700px;
		margin: 0 auto;
		padding: 2rem;
	}

	.hero {
		text-align: center;
		margin-bottom: 3rem;
		padding: 2rem 0;
		border-bottom: 2px solid #eee;
	}

	h1 {
		font-size: 2.5rem;
		margin-bottom: 0.5rem;
		color: #333;
	}

	.subtitle {
		font-size: 1.1rem;
		color: #666;
		font-weight: 300;
	}

	.form-section {
		background: #f9f9f9;
		padding: 2rem;
		border-radius: 8px;
		box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
	}

	.alert {
		padding: 1rem;
		border-radius: 6px;
		margin-bottom: 1.5rem;
		font-size: 0.95rem;
	}

	.alert-success {
		background-color: #d4edda;
		color: #155724;
		border: 1px solid #c3e6cb;
	}

	.alert-error {
		background-color: #f8d7da;
		color: #721c24;
		border: 1px solid #f5c6cb;
	}

	form {
		display: flex;
		flex-direction: column;
		gap: 1.5rem;
	}

	.form-group {
		display: flex;
		flex-direction: column;
		gap: 0.5rem;
	}

	label {
		font-weight: 600;
		color: #333;
		font-size: 0.95rem;
	}

	input,
	textarea {
		padding: 0.75rem;
		border: 1px solid #ddd;
		border-radius: 6px;
		font-size: 1rem;
		font-family: inherit;
		transition: border-color 0.2s ease, box-shadow 0.2s ease;
	}

	input:focus,
	textarea:focus {
		outline: none;
		border-color: #4a90e2;
		box-shadow: 0 0 0 3px rgba(74, 144, 226, 0.1);
	}

	input:disabled,
	textarea:disabled {
		background-color: #f5f5f5;
		cursor: not-allowed;
		opacity: 0.6;
	}

	textarea {
		resize: vertical;
		min-height: 120px;
	}

	.submit-btn {
		padding: 0.875rem 2rem;
		background-color: #4a90e2;
		color: white;
		border: none;
		border-radius: 6px;
		font-size: 1rem;
		font-weight: 600;
		cursor: pointer;
		transition: background-color 0.2s ease, transform 0.1s ease;
		align-self: flex-start;
	}

	.submit-btn:hover:not(:disabled) {
		background-color: #357abd;
		transform: translateY(-1px);
	}

	.submit-btn:active:not(:disabled) {
		transform: translateY(0);
	}

	.submit-btn:disabled {
		background-color: #94b8d8;
		cursor: not-allowed;
		opacity: 0.7;
	}

	@media (max-width: 768px) {
		.contact-container {
			padding: 1rem;
		}

		h1 {
			font-size: 2rem;
		}

		.form-section {
			padding: 1.5rem;
		}

		.submit-btn {
			width: 100%;
		}
	}
</style>
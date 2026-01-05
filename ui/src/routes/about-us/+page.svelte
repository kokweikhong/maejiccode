<script lang="ts">
	// You can add any interactive logic here
</script>

<div class="about-container">
	<section class="hero">
		<h1>About Us</h1>
		<p class="subtitle">Building innovative solutions with passion and expertise</p>
	</section>

	<section class="content">
		<div class="section">
			<h2>Our Story</h2>
			<p>
				We are a team dedicated to creating exceptional web experiences. Our journey began with a 
				simple mission: to build powerful, user-friendly applications that make a difference.
			</p>
		</div>

		<div class="section">
			<h2>What We Do</h2>
			<p>
				We specialize in developing modern web applications using cutting-edge technologies. 
				From concept to deployment, we handle every aspect of the development process with 
				attention to detail and a commitment to quality.
			</p>
		</div>

		<div class="section">
			<h2>Our Values</h2>
			<ul>
				<li><strong>Innovation:</strong> We embrace new technologies and creative solutions</li>
				<li><strong>Quality:</strong> We deliver excellence in every project</li>
				<li><strong>Collaboration:</strong> We work together to achieve the best results</li>
				<li><strong>Integrity:</strong> We build trust through honest and transparent practices</li>
			</ul>
		</div>

		<div class="section">
			<h2>Get In Touch</h2>
			<p>
				We'd love to hear from you! Whether you have a project in mind or just want to say hello, 
				feel free to reach out.
			</p>
		</div>
	</section>
</div>

<style>
	.about-container {
		max-width: 900px;
		margin: 0 auto;
		padding: 2rem;
	}

	.hero {
		text-align: center;
		margin-bottom: 3rem;
		padding: 3rem 0;
		border-bottom: 2px solid #eee;
	}

	h1 {
		font-size: 3rem;
		margin-bottom: 1rem;
		color: #333;
	}

	.subtitle {
		font-size: 1.25rem;
		color: #666;
		font-weight: 300;
	}

	.content {
		display: flex;
		flex-direction: column;
		gap: 2.5rem;
	}

	.section {
		padding: 1.5rem;
		background: #f9f9f9;
		border-radius: 8px;
		transition: transform 0.2s ease, box-shadow 0.2s ease;
	}

	.section:hover {
		transform: translateY(-2px);
		box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
	}

	h2 {
		font-size: 1.75rem;
		color: #2c3e50;
		margin-bottom: 1rem;
	}

	p {
		line-height: 1.8;
		color: #555;
		margin-bottom: 1rem;
	}

	ul {
		list-style: none;
		padding: 0;
	}

	li {
		padding: 0.75rem 0;
		color: #555;
		line-height: 1.6;
	}

	strong {
		color: #2c3e50;
	}

	@media (max-width: 768px) {
		.about-container {
			padding: 1rem;
		}

		h1 {
			font-size: 2rem;
		}

		.subtitle {
			font-size: 1rem;
		}

		.section {
			padding: 1rem;
		}
	}
</style>
<script lang="ts">
	interface Service {
		title: string;
		description: string;
		icon: string;
		gradient: string;
	}
	
	interface Props {
		mouseX: number;
        mouseY: number;
		services: Service[];
	}
	
	let { mouseX, mouseY, services }: Props = $props();
</script>

<section id="services" class="bg-linear-to-b from-white to-gray-50 py-20 relative overflow-hidden">
	<!-- Decorative elements -->
	<div class="absolute top-20 left-10 w-72 h-72 bg-purple-300 rounded-full mix-blend-multiply filter blur-xl opacity-20 animate-blob"></div>
	<div class="absolute top-40 right-10 w-72 h-72 bg-yellow-300 rounded-full mix-blend-multiply filter blur-xl opacity-20 animate-blob animation-delay-2000"></div>
	<div class="absolute bottom-20 left-1/2 w-72 h-72 bg-pink-300 rounded-full mix-blend-multiply filter blur-xl opacity-20 animate-blob animation-delay-4000"></div>
	
	<!-- 3D Geometric Shapes -->
	<div class="absolute inset-0 overflow-hidden pointer-events-none">
		<!-- Floating Torus -->
		<div class="torus-container" style="transform: translate3d({mouseX * 40}px, {mouseY * 40}px, 0)">
			<div class="torus"></div>
		</div>
		<!-- Floating Octahedron -->
		<div class="octahedron-container" style="transform: translate3d({mouseX * -35}px, {mouseY * -35}px, 0)">
			<div class="octahedron">
				<div class="octa-face"></div>
				<div class="octa-face"></div>
				<div class="octa-face"></div>
				<div class="octa-face"></div>
				<div class="octa-face"></div>
				<div class="octa-face"></div>
				<div class="octa-face"></div>
				<div class="octa-face"></div>
			</div>
		</div>
	</div>
	
	<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 relative z-10">
		<div class="text-center mb-16">
			<h2 class="text-3xl md:text-5xl font-bold text-gray-900 mb-4 animate-text-shimmer">
				Our Services
			</h2>
			<p class="text-lg text-gray-600 max-w-2xl mx-auto">
				Comprehensive web development services to power your digital success
			</p>
		</div>
		
		<div class="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
			{#each services as service, i}
				<div 
					class="group relative bg-white p-8 rounded-2xl shadow-xl hover:shadow-2xl transition-all duration-500 transform hover:-translate-y-3 card-3d overflow-hidden"
					style="animation: fadeInUp 0.6s ease-out {i * 0.1}s backwards"
				>
					<!-- Gradient overlay -->
				<div class="absolute inset-0 bg-linear-to-br {service.gradient} opacity-0 group-hover:opacity-10 transition-opacity duration-500"></div>
					<div class="relative">
						<div class="text-6xl mb-4 transform group-hover:scale-125 group-hover:rotate-12 transition-all duration-500 filter group-hover:drop-shadow-2xl animate-float" style="animation-delay: {i * 0.1}s">
							{service.icon}
						</div>
						<div class="absolute top-0 left-0 text-6xl opacity-30 blur-sm transform translate-x-2 translate-y-2">
							{service.icon}
						</div>
					</div>
					
					<h3 class="text-xl font-bold text-gray-900 mb-3 group-hover:text-transparent group-hover:bg-clip-text group-hover:bg-linear-to-r group-hover:{service.gradient} transition-all duration-300">
						{service.title}
					</h3>
					<p class="text-gray-600 leading-relaxed mb-4">
						{service.description}
					</p>
					
					<!-- Animated border -->
					<div class="absolute inset-0 rounded-2xl border-2 border-transparent group-hover:border-linear-to-r group-hover:{service.gradient} opacity-0 group-hover:opacity-100 transition-opacity duration-500 animate-border-flow"></div>
					
					<!-- Shine effect -->
					<div class="absolute inset-0 rounded-2xl overflow-hidden">
						<div class="shine"></div>
					</div>
				</div>
			{/each}
		</div>
	</div>
</section>

<style>
	/* Optimize 3D shapes for mobile */
	@media (max-width: 768px) {
		.torus-container {
			width: 90px;
			height: 90px;
			top: 5%;
			right: 5%;
			opacity: 0.4;
		}
		
		.octahedron-container {
			width: 80px;
			height: 80px;
			bottom: 10%;
			left: 5%;
			opacity: 0.4;
		}
	}
</style>

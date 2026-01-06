<script lang="ts">
	interface Props {
		mouseX: number;
		mouseY: number;
		scrollY: number;
	}
	
	let { mouseX, mouseY, scrollY }: Props = $props();
</script>

<section class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-20 relative">
	<!-- 3D Decorative Shapes -->
	<div class="absolute inset-0 overflow-hidden pointer-events-none">
		<!-- Floating Pyramid Left -->
		<div class="pyramid-container left" style="transform: translate3d({mouseX * 30}px, {mouseY * 30}px, 0)">
			<div class="pyramid">
				<div class="pyramid-face front"></div>
				<div class="pyramid-face back"></div>
				<div class="pyramid-face right"></div>
				<div class="pyramid-face left"></div>
				<div class="pyramid-face bottom"></div>
			</div>
		</div>
		<!-- Floating Sphere Right -->
		<div class="sphere-container right" style="transform: translate3d({mouseX * -25}px, {mouseY * -25}px, 0)">
			<div class="sphere"></div>
		</div>
	</div>
	
	<div class="text-center mb-16" style="transform: translateY({scrollY * -0.1}px)">
		<h2 class="text-3xl md:text-5xl font-bold text-gray-900 mb-4 animate-text-shimmer">
			Why Choose Maejic Code Studio?
		</h2>
		<p class="text-lg text-gray-600 max-w-3xl mx-auto">
			We are a team of passionate developers dedicated to delivering exceptional web solutions. 
			With expertise across modern technologies and a commitment to quality, we turn your vision into reality.
		</p>
	</div>
	
	<div class="grid md:grid-cols-3 gap-8 mb-16">
		<div class="group text-center p-8 bg-white rounded-2xl shadow-xl hover:shadow-2xl transition-all duration-500 transform hover:-translate-y-4 hover:rotate-1 card-3d">
			<div class="text-6xl mb-4 transform group-hover:scale-125 group-hover:rotate-12 transition-all duration-500 animate-float">⚡</div>
			<h3 class="text-2xl font-bold text-gray-900 mb-3 group-hover:text-blue-600 transition-colors">Fast & Efficient</h3>
			<p class="text-gray-600">Lightning-fast websites optimized for performance and user experience</p>
			<div class="mt-4 h-1 w-0 group-hover:w-full bg-linear-to-r from-blue-500 to-purple-500 transition-all duration-500 mx-auto"></div>
		</div>
		<div class="group text-center p-8 bg-white rounded-2xl shadow-xl hover:shadow-2xl transition-all duration-500 transform hover:-translate-y-4 hover:rotate-1 card-3d">
			<div class="text-6xl mb-4 transform group-hover:scale-125 group-hover:rotate-12 transition-all duration-500 animate-float" style="animation-delay: 0.2s">🎯</div>
			<h3 class="text-2xl font-bold text-gray-900 mb-3 group-hover:text-purple-600 transition-colors">Result-Driven</h3>
			<p class="text-gray-600">Solutions focused on achieving your business goals and ROI</p>
			<div class="mt-4 h-1 w-0 group-hover:w-full bg-linear-to-r from-purple-500 to-pink-500 transition-all duration-500 mx-auto"></div>
		</div>
		<div class="group text-center p-8 bg-white rounded-2xl shadow-xl hover:shadow-2xl transition-all duration-500 transform hover:-translate-y-4 hover:rotate-1 card-3d">
			<div class="text-6xl mb-4 transform group-hover:scale-125 group-hover:rotate-12 transition-all duration-500 animate-float" style="animation-delay: 0.4s">🤝</div>
			<h3 class="text-2xl font-bold text-gray-900 mb-3 group-hover:text-indigo-600 transition-colors">Reliable Partner</h3>
			<p class="text-gray-600">Ongoing support and maintenance to keep your site running smoothly</p>
			<div class="mt-4 h-1 w-0 group-hover:w-full bg-linear-to-r from-indigo-500 to-blue-500 transition-all duration-500 mx-auto"></div>
		</div>
	</div>
</section>

<style>
	/* Optimize 3D shapes for mobile */
	@media (max-width: 768px) {
		.pyramid-container {
			width: 100px;
			height: 100px;
			top: 5%;
			left: 2%;
			opacity: 0.4;
		}
		
		.sphere-container {
			width: 80px;
			height: 80px;
			top: 15%;
			right: 2%;
			opacity: 0.4;
		}
	}
</style>

<script lang="ts">
	interface Props {
		mouseX: number;
		mouseY: number;
		isVisible: boolean;
	}
	
	let { mouseX, mouseY, isVisible }: Props = $props();
</script>

<section class="relative overflow-hidden bg-linear-to-r from-blue-600 via-purple-600 to-indigo-700 text-white">
	<!-- Animated gradient overlay -->
	<div class="absolute inset-0 bg-linear-to-r from-blue-600/50 via-purple-600/50 to-pink-600/50 animate-gradient"></div>
	
	<!-- 3D Geometric shapes -->
	<div class="absolute inset-0 overflow-hidden pointer-events-none">
		<div class="cube-container" style="transform: translate3d({mouseX * 50}px, {mouseY * 50}px, 0)">
			<div class="cube">
				<div class="cube-face front"></div>
				<div class="cube-face back"></div>
				<div class="cube-face right"></div>
				<div class="cube-face left"></div>
				<div class="cube-face top"></div>
				<div class="cube-face bottom"></div>
			</div>
		</div>
		<div class="cube-container-2" style="transform: translate3d({mouseX * -30}px, {mouseY * -30}px, 0)">
			<div class="cube-2">
				<div class="cube-face front"></div>
				<div class="cube-face back"></div>
				<div class="cube-face right"></div>
				<div class="cube-face left"></div>
				<div class="cube-face top"></div>
				<div class="cube-face bottom"></div>
			</div>
		</div>
	</div>
	
	<div class="relative z-10 max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-24 md:py-32">
		<div class="text-center">
			<div class="mb-8 inline-block">
				<div class="text-8xl mb-4 animate-bounce-slow">✨</div>
			</div>
			<h1 class="text-4xl md:text-7xl font-bold mb-6 leading-tight animate-text-glow">
				Welcome to <span class="text-transparent bg-clip-text bg-linear-to-r from-yellow-300 via-pink-300 to-purple-300 animate-gradient-text">Maejic Code Studio</span>
			</h1>
			<p class="text-xl md:text-3xl mb-8 text-blue-100 max-w-3xl mx-auto font-light animate-pulse-slow">
				Transforming Ideas Into Digital Excellence
			</p>
			<p class="text-lg md:text-xl mb-12 text-blue-50 max-w-2xl mx-auto">
				We craft innovative web solutions that elevate your business and engage your audience
			</p>
			<div class="flex flex-col sm:flex-row gap-4 justify-center">
				<a href="/contact-us" class="group relative bg-white text-blue-600 px-8 py-4 rounded-lg font-semibold text-lg overflow-hidden transition-all transform hover:scale-105 shadow-2xl">
					<span class="relative z-10">Get Started</span>
					<div class="absolute inset-0 bg-linear-to-r from-blue-600 to-purple-600 transform scale-x-0 group-hover:scale-x-100 transition-transform origin-left"></div>
					<span class="absolute inset-0 flex items-center justify-center text-white opacity-0 group-hover:opacity-100 transition-opacity">Get Started</span>
				</a>
				<a href="#services" class="group relative border-2 border-white text-white px-8 py-4 rounded-lg font-semibold text-lg transition-all transform hover:scale-105 perspective-1000">
					<span class="relative z-10 group-hover:transform group-hover:rotateX-12">Our Services</span>
				</a>
			</div>
		</div>
	</div>
	<div class="absolute bottom-0 left-0 right-0">
		<svg viewBox="0 0 1440 120" class="w-full h-auto">
			<path fill="#f8fafc" d="M0,64L80,69.3C160,75,320,85,480,80C640,75,800,53,960,48C1120,43,1280,53,1360,58.7L1440,64L1440,120L1360,120C1280,120,1120,120,960,120C800,120,640,120,480,120C320,120,160,120,80,120L0,120Z"></path>
		</svg>
	</div>
</section>

<style>
	/* Optimize 3D shapes for mobile */
	@media (max-width: 768px) {
		.cube-container {
			width: 100px;
			height: 100px;
			top: 5%;
			right: 5%;
			opacity: 0.5;
		}
		
		.cube-container-2 {
			width: 80px;
			height: 80px;
			bottom: 10%;
			left: 5%;
			opacity: 0.5;
		}
		
		.cube-face {
			width: 100px !important;
			height: 100px !important;
		}
		
		.cube-2 .cube-face {
			width: 80px !important;
			height: 80px !important;
		}
	}
</style>

<script lang="ts">
  import { onMount } from 'svelte';
  import { services } from './constants';
  import HeroSection from './_components/HeroSection.svelte';
  import AboutSection from './_components/AboutSection.svelte';
  import ServicesSection from './_components/ServicesSection.svelte';

  let isVisible = $state(false);
  let mouseX = $state(0);
  let mouseY = $state(0);
  let scrollY = $state(0);

  onMount(() => {
    isVisible = true;

    const handleMouseMove = (e: MouseEvent) => {
      mouseX = e.clientX / window.innerWidth;
      mouseY = e.clientY / window.innerHeight;
    };

    const handleScroll = () => {
      scrollY = window.scrollY;
    };

    window.addEventListener('mousemove', handleMouseMove);
    window.addEventListener('scroll', handleScroll);

    return () => {
      window.removeEventListener('mousemove', handleMouseMove);
      window.removeEventListener('scroll', handleScroll);
    };
  });

  const floatingShapes = Array.from({ length: 20 }, (_, i) => ({
    id: i,
    size: Math.random() * 100 + 50,
    left: Math.random() * 100,
    top: Math.random() * 100,
    duration: Math.random() * 20 + 10,
    delay: Math.random() * 5,
  }));
</script>

<svelte:head>
  <title>Maejic Code Studio - Professional Web Development Services</title>
  <meta
    name="description"
    content="Maejic Code Studio offers custom website development, web applications, e-commerce solutions, and more. Transform your digital presence with our expert team."
  />
</svelte:head>

<div
  class="min-h-screen bg-linear-to-br from-slate-50 to-blue-50 relative overflow-hidden"
>
  <!-- Animated Background Shapes -->
  <div class="fixed inset-0 pointer-events-none overflow-hidden">
    {#each floatingShapes as shape}
      <div
        class="absolute rounded-full bg-linear-to-br from-blue-400/20 to-purple-400/20 blur-3xl animate-float"
        style="
					width: {shape.size}px;
					height: {shape.size}px;
					left: {shape.left}%;
					top: {shape.top}%;
					animation-duration: {shape.duration}s;
					animation-delay: {shape.delay}s;
				"
      ></div>
    {/each}
  </div>

  <!-- Hero Section -->
  <HeroSection {mouseX} {mouseY} {isVisible} />

  <!-- About Section -->
  <AboutSection {mouseX} {mouseY} {scrollY} />

  <!-- Services Section -->
  <ServicesSection {mouseX} {mouseY} {services} />

  <!-- CTA Section -->
  <section
    class="relative bg-linear-to-r from-blue-600 via-purple-600 to-indigo-700 text-white py-20 overflow-hidden"
  >
    <!-- Animated background -->
    <div class="absolute inset-0 opacity-30">
      <div
        class="absolute inset-0 bg-linear-to-r from-blue-400 to-purple-400 animate-gradient"
      ></div>
    </div>

    <!-- 3D floating elements -->
    <div class="absolute inset-0 overflow-hidden">
      <div class="floating-orb orb-1"></div>
      <div class="floating-orb orb-2"></div>
      <div class="floating-orb orb-3"></div>
    </div>

    <div
      class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 text-center relative z-10"
    >
      <div class="text-7xl mb-6 animate-bounce-slow inline-block">🚀</div>
      <h2 class="text-3xl md:text-5xl font-bold mb-6 animate-text-glow">
        Ready to Start Your Project?
      </h2>
      <p class="text-xl md:text-2xl mb-8 text-blue-100">
        Let's discuss how we can help bring your vision to life
      </p>
      <div class="flex flex-col sm:flex-row gap-4 justify-center">
        <a
          href="/contact-us"
          class="group relative bg-white text-blue-600 px-8 py-4 rounded-xl font-semibold text-lg overflow-hidden transition-all transform hover:scale-110 shadow-2xl hover:shadow-3xl"
        >
          <span class="relative z-10">Contact Us Today</span>
          <div
            class="absolute inset-0 bg-linear-to-r from-yellow-400 to-pink-500 transform scale-x-0 group-hover:scale-x-100 transition-transform origin-left"
          ></div>
          <span
            class="absolute inset-0 flex items-center justify-center text-white opacity-0 group-hover:opacity-100 transition-opacity font-bold"
            >Contact Us Today</span
          >
        </a>
        <a
          href="/about-us"
          class="group relative border-2 border-white text-white px-8 py-4 rounded-xl font-semibold text-lg transition-all transform hover:scale-110 backdrop-blur-sm bg-white/10 hover:bg-white hover:text-purple-600 shadow-2xl"
        >
          Learn More About Us
        </a>
      </div>
    </div>
  </section>

  <!-- Footer -->
  <footer class="bg-gray-900 text-gray-300 py-12">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="grid md:grid-cols-3 gap-8 mb-8">
        <div>
          <h3 class="text-white text-xl font-bold mb-4">Maejic Code Studio</h3>
          <p class="text-gray-400">
            Professional web development services to elevate your digital
            presence
          </p>
        </div>
        <div>
          <h3 class="text-white text-lg font-semibold mb-4">Quick Links</h3>
          <ul class="space-y-2">
            <li>
              <a href="/" class="hover:text-white transition-colors">Home</a>
            </li>
            <li>
              <a href="/about-us" class="hover:text-white transition-colors"
                >About Us</a
              >
            </li>
            <li>
              <a href="/contact-us" class="hover:text-white transition-colors"
                >Contact</a
              >
            </li>
          </ul>
        </div>
        <div>
          <h3 class="text-white text-lg font-semibold mb-4">Contact</h3>
          <p class="text-gray-400">
            <a
              href="https://maejiccode.com"
              target="_blank"
              rel="noopener noreferrer"
              class="hover:text-white transition-colors"
            >
              www.maejiccode.com
            </a>
          </p>
        </div>
      </div>
      <div class="border-t border-gray-800 pt-8 text-center text-gray-400">
        <p>
          &copy; {new Date().getFullYear()} Maejic Code Studio. All rights reserved.
        </p>
      </div>
    </div>
  </footer>
</div>

<style>
  :global(html) {
    scroll-behavior: smooth;
  }

  /* Floating Orbs */
  .floating-orb {
    position: absolute;
    border-radius: 50%;
    background: radial-gradient(
      circle at 30% 30%,
      rgba(255, 255, 255, 0.8),
      rgba(255, 255, 255, 0.1)
    );
    filter: blur(40px);
    animation: float 8s ease-in-out infinite;
  }

  .orb-1 {
    width: 300px;
    height: 300px;
    top: 10%;
    left: 10%;
    animation-delay: 0s;
  }

  .orb-2 {
    width: 200px;
    height: 200px;
    top: 50%;
    right: 15%;
    animation-delay: 2s;
  }

  .orb-3 {
    width: 250px;
    height: 250px;
    bottom: 10%;
    left: 50%;
    animation-delay: 4s;
  }

  /* Animations */
  @keyframes float {
    0%,
    100% {
      transform: translateY(0px) translateX(0px);
    }
    25% {
      transform: translateY(-20px) translateX(10px);
    }
    50% {
      transform: translateY(-40px) translateX(-10px);
    }
    75% {
      transform: translateY(-20px) translateX(10px);
    }
  }

  @keyframes rotate3d {
    0% {
      transform: rotateX(0deg) rotateY(0deg);
    }
    100% {
      transform: rotateX(360deg) rotateY(360deg);
    }
  }

  @keyframes gradient {
    0%,
    100% {
      background-position: 0% 50%;
    }
    50% {
      background-position: 100% 50%;
    }
  }

  @keyframes blob {
    0%,
    100% {
      transform: translate(0, 0) scale(1);
    }
    25% {
      transform: translate(20px, -50px) scale(1.1);
    }
    50% {
      transform: translate(-20px, 20px) scale(0.9);
    }
    75% {
      transform: translate(50px, 50px) scale(1.05);
    }
  }

  @keyframes fadeInUp {
    from {
      opacity: 0;
      transform: translateY(30px);
    }
    to {
      opacity: 1;
      transform: translateY(0);
    }
  }

  @keyframes wave {
    0%,
    100% {
      transform: translateX(0);
    }
    50% {
      transform: translateX(-25px);
    }
  }

  @keyframes textGlow {
    0%,
    100% {
      text-shadow:
        0 0 20px rgba(255, 255, 255, 0.5),
        0 0 40px rgba(255, 255, 255, 0.3);
    }
    50% {
      text-shadow:
        0 0 30px rgba(255, 255, 255, 0.8),
        0 0 60px rgba(255, 255, 255, 0.5);
    }
  }

  @keyframes shimmer {
    0% {
      background-position: -1000px 0;
    }
    100% {
      background-position: 1000px 0;
    }
  }

  .animate-gradient {
    background-size: 200% 200%;
    animation: gradient 15s ease infinite;
  }

  .animate-text-glow {
    animation: textGlow 3s ease-in-out infinite;
  }

  .animate-bounce-slow {
    animation: bounce 3s infinite;
  }

  @keyframes bounce {
    0%,
    100% {
      transform: translateY(0);
    }
    50% {
      transform: translateY(-20px);
    }
  }

  @keyframes pulse {
    0%,
    100% {
      opacity: 1;
    }
    50% {
      opacity: 0.8;
    }
  }

  .animate-float {
    animation: float 6s ease-in-out infinite;
  }
</style>

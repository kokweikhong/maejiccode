<script lang="ts">
  import { onMount } from 'svelte';

  const navLinks: { name: string; href: string }[] = [
    { name: 'Home', href: '/' },
    { name: 'Services', href: '#services' },
    { name: 'About Us', href: '/about-us' },
    { name: 'Contact Us', href: '/contact-us' },
    { name: 'Games', href: '/games' },
  ];

  let mobileMenuOpen = $state(false);
  let navbarScrolled = $state(false);

  onMount(() => {
    // Set initial state based on current scroll position
    navbarScrolled = window.scrollY > 150;

    const handleScroll = () => {
      navbarScrolled = window.scrollY > 150;
    };

    window.addEventListener('scroll', handleScroll);

    return () => {
      window.removeEventListener('scroll', handleScroll);
    };
  });
</script>

<nav
  class="fixed top-0 left-0 right-0 z-50 transition-all duration-300 {navbarScrolled
    ? 'bg-white/95 backdrop-blur-lg shadow-lg'
    : 'bg-black/10 backdrop-blur-sm'}"
>
  <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
    <div class="flex justify-between items-center h-20">
      <!-- Logo -->
      <a href="/" class="flex items-center space-x-2 group">
        <div
          class="text-3xl transform group-hover:rotate-12 transition-transform duration-300"
        >
          ✨
        </div>
        <span
          class="text-2xl font-bold {navbarScrolled
            ? 'text-transparent bg-clip-text bg-linear-to-r from-blue-600 to-purple-600'
            : 'text-white drop-shadow-lg'} transition-colors"
        >
          Maejic Code
        </span>
      </a>

      <!-- Desktop Navigation -->
      <div class="hidden md:flex items-center space-x-8">
        {#each navLinks as link}
          <a
            href={link.href}
            class="px-3 py-2 rounded-lg font-semibold transition-colors {navbarScrolled
              ? 'text-gray-700 hover:bg-gray-100'
              : 'text-white hover:bg-white/10'}"
          >
            {link.name}
          </a>
        {/each}
        <a
          href="/contact-us"
          class="px-6 py-2.5 rounded-lg font-semibold transition-all transform hover:scale-105 shadow-lg {navbarScrolled
            ? 'bg-linear-to-r from-blue-600 to-purple-600 text-white hover:shadow-xl'
            : 'bg-white text-blue-600 hover:bg-yellow-300 hover:text-blue-700'}"
        >
          Contact Us
        </a>
      </div>

      <!-- Mobile Menu Button -->
      <button
        onclick={() => (mobileMenuOpen = !mobileMenuOpen)}
        class="md:hidden p-2 rounded-lg transition-colors {navbarScrolled
          ? 'text-gray-700 hover:bg-gray-100'
          : 'text-white hover:bg-white/10'}"
        aria-label="Toggle menu"
      >
        <svg
          class="w-6 h-6"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
        >
          {#if mobileMenuOpen}
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M6 18L18 6M6 6l12 12"
            ></path>
          {:else}
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M4 6h16M4 12h16M4 18h16"
            ></path>
          {/if}
        </svg>
      </button>
    </div>
  </div>

  <!-- Mobile Menu -->
  {#if mobileMenuOpen}
    <div
      class="md:hidden bg-white border-t border-gray-200 shadow-lg animate-slide-down"
    >
      <div class="px-4 py-4 space-y-3">
        {#each navLinks as link}
          <a
            href={link.href}
            class="block px-4 py-3 rounded-lg font-semibold {navbarScrolled
              ? 'text-gray-700 hover:bg-gray-100'
              : 'text-black hover:bg-gray-100'} transition-colors"
            onclick={() => (mobileMenuOpen = false)}
          >
            {link.name}
          </a>
        {/each}
        <a
          href="/contact-us"
          class="block px-4 py-3 bg-linear-to-r from-blue-600 to-purple-600 text-white rounded-lg font-semibold text-center hover:shadow-lg transition-all"
          onclick={() => (mobileMenuOpen = false)}
        >
          📧 Contact Us
        </a>
      </div>
    </div>
  {/if}
</nav>

<style>
  /* Mobile menu animation */
  @keyframes slideDown {
    from {
      opacity: 0;
      transform: translateY(-10px);
    }
    to {
      opacity: 1;
      transform: translateY(0);
    }
  }

  .animate-slide-down {
    animation: slideDown 0.3s ease-out;
  }
</style>

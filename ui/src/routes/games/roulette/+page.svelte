<script lang="ts">
  import { onMount } from 'svelte';
  import Wheel from './_components/Wheel.svelte';
  import Nameslist from './_components/Nameslist.svelte';

  let canvas: HTMLCanvasElement | undefined;
  let ctx: CanvasRenderingContext2D;

  // Game state
  let names = $state<string[]>([]);
  let currentName = $state('');
  let isSpinning = $state(false);
  let rotation = $state(0);
  let targetRotation = $state(0);
  let spinSpeed = $state(0);
  let winner = $state('');
  let showWinner = $state(false);

  // Colors for wheel segments
  const colors = [
    '#FF6B6B',
    '#4ECDC4',
    '#45B7D1',
    '#FFA07A',
    '#98D8C8',
    '#F7DC6F',
    '#BB8FCE',
    '#85C1E2',
    '#F8B739',
    '#52B788',
    '#E76F51',
    '#2A9D8F',
  ];

  function handleCanvasReady(canvasElement: HTMLCanvasElement) {
    canvas = canvasElement;
    initializeCanvas();
  }

  function initializeCanvas() {
    if (!canvas) return;
    ctx = canvas.getContext('2d')!;

    // Set responsive canvas size
    const updateCanvasSize = () => {
      if (!canvas) return;
      const size = Math.min(
        500,
        window.innerWidth - 40,
        window.innerHeight - 300,
      );
      canvas.width = size;
      canvas.height = size;
    };

    updateCanvasSize();
    window.addEventListener('resize', updateCanvasSize);

    // Load names from localStorage
    const savedNames = localStorage.getItem('rouletteNames');
    if (savedNames) {
      try {
        names = JSON.parse(savedNames);
      } catch (e) {
        console.error('Failed to load saved names:', e);
      }
    }

    const animate = () => {
      drawWheel();

      if (isSpinning) {
        rotation += spinSpeed;
        spinSpeed *= 0.98; // Deceleration

        if (spinSpeed < 0.1) {
          isSpinning = false;
          spinSpeed = 0;
          rotation = targetRotation;
          determineWinner();
        }
      }

      requestAnimationFrame(animate);
    };

    animate();

    return () => {
      window.removeEventListener('resize', updateCanvasSize);
    };
  }

  function startSpin() {
    if (names.length < 2 || isSpinning) return;

    showWinner = false;
    winner = '';

    // Random number of full rotations (5-10) plus random position
    const fullRotations = 5 + Math.floor(Math.random() * 5);
    const randomAngle = Math.random() * 360;
    targetRotation = (fullRotations * 360 + randomAngle) % 360;

    spinSpeed = 20 + Math.random() * 10;
    isSpinning = true;
  }

  onMount(() => {
    // Load names from localStorage
    const savedNames = localStorage.getItem('rouletteNames');
    if (savedNames) {
      try {
        names = JSON.parse(savedNames);
      } catch (e) {
        console.error('Failed to load saved names:', e);
      }
    }
  });

  function addName() {
    if (currentName.trim() && names.length < 20) {
      names = [...names, currentName.trim()];
      currentName = '';
      // Save to localStorage
      localStorage.setItem('rouletteNames', JSON.stringify(names));
    }
  }

  function removeName(index: number) {
    names = names.filter((_, i) => i !== index);
    // Save to localStorage
    localStorage.setItem('rouletteNames', JSON.stringify(names));
  }

  function clearAllNames() {
    if (confirm('Are you sure you want to clear all names?')) {
      names = [];
      localStorage.removeItem('rouletteNames');
    }
  }

  function determineWinner() {
    if (names.length === 0) return;

    // Calculate which segment the pointer is on
    const segmentAngle = 360 / names.length;
    const adjustedRotation = (rotation + 90) % 360; // Adjust for pointer at top
    const winnerIndex =
      Math.floor((360 - adjustedRotation) / segmentAngle) % names.length;

    winner = names[winnerIndex];
    showWinner = true;
  }

  function drawWheel() {
    if (!ctx || !canvas) return;

    ctx.clearRect(0, 0, canvas.width, canvas.height);

    const centerX = canvas.width / 2;
    const centerY = canvas.height / 2;
    const radius = Math.min(canvas.width, canvas.height) * 0.4; // Responsive radius

    if (names.length === 0) {
      // Draw empty wheel
      ctx.fillStyle = '#e0e0e0';
      ctx.beginPath();
      ctx.arc(centerX, centerY, radius, 0, Math.PI * 2);
      ctx.fill();

      ctx.fillStyle = '#666';
      ctx.font = `${radius * 0.1}px Inter, sans-serif`;
      ctx.textAlign = 'center';
      ctx.textBaseline = 'middle';
      ctx.fillText('Add names to start', centerX, centerY);
      return;
    }

    const segmentAngle = (Math.PI * 2) / names.length;

    // Draw wheel segments
    names.forEach((name, i) => {
      const startAngle = (rotation * Math.PI) / 180 + i * segmentAngle;
      const endAngle = startAngle + segmentAngle;

      // Draw segment
      ctx.fillStyle = colors[i % colors.length];
      ctx.beginPath();
      ctx.moveTo(centerX, centerY);
      ctx.arc(centerX, centerY, radius, startAngle, endAngle);
      ctx.closePath();
      ctx.fill();

      // Draw border
      ctx.strokeStyle = '#fff';
      ctx.lineWidth = 3;
      ctx.stroke();

      // Draw text
      ctx.save();
      ctx.translate(centerX, centerY);
      ctx.rotate(startAngle + segmentAngle / 2);
      ctx.textAlign = 'right';
      ctx.textBaseline = 'middle';
      ctx.fillStyle = '#fff';
      ctx.font = `bold ${Math.max(12, radius * 0.08)}px Inter, sans-serif`;
      ctx.shadowColor = 'rgba(0, 0, 0, 0.5)';
      ctx.shadowBlur = 3;
      ctx.fillText(name, radius - radius * 0.1, 0);
      ctx.restore();
    });

    // Draw center circle
    ctx.fillStyle = '#fff';
    ctx.beginPath();
    ctx.arc(centerX, centerY, radius * 0.15, 0, Math.PI * 2);
    ctx.fill();
    ctx.strokeStyle = '#333';
    ctx.lineWidth = 3;
    ctx.stroke();

    // Draw center text
    ctx.fillStyle = '#333';
    ctx.font = `bold ${Math.max(10, radius * 0.07)}px Inter, sans-serif`;
    ctx.textAlign = 'center';
    ctx.textBaseline = 'middle';
    ctx.shadowColor = 'transparent';
    ctx.fillText('SPIN', centerX, centerY);

    // Draw pointer (triangle at top) - responsive size
    const pointerSize = radius * 0.075; // Relative to radius
    ctx.fillStyle = '#FF1744';
    ctx.beginPath();
    ctx.moveTo(centerX, centerY - radius - pointerSize);
    ctx.lineTo(centerX - pointerSize, centerY - radius - pointerSize * 2);
    ctx.lineTo(centerX + pointerSize, centerY - radius - pointerSize * 2);
    ctx.closePath();
    ctx.fill();
    ctx.strokeStyle = '#fff';
    ctx.lineWidth = 2;
    ctx.stroke();
  }

  function handleKeyPress(e: KeyboardEvent) {
    if (e.key === 'Enter') {
      addName();
    }
  }
</script>

<svelte:head>
  <title>Roulette Game | Maejic Code Studio</title>
</svelte:head>

<div
  class="min-h-screen bg-linear-to-br from-purple-900 via-blue-900 to-indigo-900 pb-12 pt-30 px-4 sm:px-6 lg:px-8"
>
  <div class="max-w-6xl mx-auto px-4">
    <!-- Header -->
    <div class="text-center mb-6 md:mb-8">
      <h1
        class="text-3xl sm:text-4xl md:text-5xl lg:text-6xl font-bold text-white mb-3 md:mb-4"
      >
        🎰 Name Roulette
      </h1>
      <p class="text-base sm:text-lg md:text-xl text-gray-300">
        Add names and spin the wheel to pick a winner!
      </p>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 md:gap-8">
      <!-- Left side - Wheel -->
      <Wheel
        {names}
        {isSpinning}
        {startSpin}
        onCanvasReady={handleCanvasReady}
      />

      <!-- Right side - Names list -->
      <Nameslist
        {names}
        {isSpinning}
        bind:currentName
        {addName}
        {removeName}
        {clearAllNames}
      />
    </div>

    <!-- Winner announcement -->
    {#if showWinner}
      <div
        class="fixed inset-0 flex items-center justify-center bg-black/80 backdrop-blur-sm z-50 animate-fadeIn p-4"
      >
        <div
          class="bg-white rounded-3xl p-6 sm:p-8 md:p-12 max-w-2xl w-full mx-4 text-center transform animate-scaleIn"
        >
          <div
            class="text-5xl sm:text-6xl md:text-8xl mb-4 md:mb-6 animate-bounce"
          >
            🎉
          </div>
          <h2
            class="text-2xl sm:text-3xl md:text-4xl lg:text-6xl font-bold text-gray-900 mb-3 md:mb-4"
          >
            Winner!
          </h2>
          <div
            class="text-3xl sm:text-4xl md:text-5xl lg:text-7xl font-bold bg-linear-to-r from-pink-500 to-purple-600 bg-clip-text text-transparent mb-6 md:mb-8 wrap-break-word"
          >
            {winner}
          </div>
          <button
            onclick={() => (showWinner = false)}
            class="bg-linear-to-r from-blue-500 to-purple-600 text-white px-6 sm:px-8 py-3 sm:py-4 rounded-xl font-semibold text-lg sm:text-xl hover:scale-105 transition-transform"
          >
            Close
          </button>
        </div>
      </div>
    {/if}

    <!-- Back button -->
    <div class="mt-8 text-center">
      <a
        href="/games"
        class="inline-block bg-white/90 hover:bg-white text-gray-900 px-6 py-3 rounded-lg font-semibold transition-all transform hover:scale-105"
      >
        ← Back to Games
      </a>
    </div>
  </div>
</div>

<style>
  @keyframes fadeIn {
    from {
      opacity: 0;
    }
    to {
      opacity: 1;
    }
  }

  @keyframes scaleIn {
    from {
      transform: scale(0.5);
      opacity: 0;
    }
    to {
      transform: scale(1);
      opacity: 1;
    }
  }

  .animate-fadeIn {
    animation: fadeIn 0.3s ease-out;
  }

  .animate-scaleIn {
    animation: scaleIn 0.5s ease-out;
  }
</style>

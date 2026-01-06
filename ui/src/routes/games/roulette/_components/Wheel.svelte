<script lang="ts">
  import { onMount } from 'svelte';

  interface Props {
    names: string[];
    isSpinning: boolean;
    startSpin: () => void;
    onCanvasReady?: (canvas: HTMLCanvasElement) => void;
  }

  let { names, isSpinning, startSpin, onCanvasReady }: Props = $props();
  let canvas: HTMLCanvasElement;

  onMount(() => {
    if (canvas && onCanvasReady) {
      onCanvasReady(canvas);
    }
  });
</script>

<div class="bg-white/10 backdrop-blur-sm rounded-2xl p-4 md:p-6">
  <h2 class="text-xl md:text-2xl font-bold text-white mb-4 text-center">
    Roulette Wheel
  </h2>

  <div class="flex justify-center mb-4 md:mb-6">
    <canvas bind:this={canvas} class="rounded-lg shadow-2xl max-w-full"
    ></canvas>
  </div>

  <div class="text-center">
    <button
      onclick={startSpin}
      disabled={names.length < 2 || isSpinning}
      class="bg-linear-to-r from-pink-500 to-purple-600 text-white px-6 sm:px-8 md:px-12 py-3 md:py-4 rounded-xl font-bold text-lg sm:text-xl md:text-2xl disabled:opacity-50 disabled:cursor-not-allowed hover:scale-105 transition-transform shadow-2xl cursor-pointer"
    >
      {isSpinning ? '🎲 SPINNING...' : '🎲 SPIN'}
    </button>

    {#if names.length < 2}
      <p class="text-yellow-300 mt-3 text-xs sm:text-sm">
        Add at least 2 names to spin
      </p>
    {/if}
  </div>
</div>

<style>
  canvas {
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    max-width: 100%;
    height: auto;
  }
</style>

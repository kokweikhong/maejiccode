<script lang="ts">
  interface Props {
    names: string[];
    isSpinning: boolean;
    currentName: string;
    addName: () => void;
    removeName: (index: number) => void;
    clearAllNames: () => void;
  }

  function handleKeyPress(e: KeyboardEvent) {
    if (e.key === 'Enter') {
      addName();
    }
  }

  let {
    names,
    isSpinning,
    currentName = $bindable(),
    addName,
    removeName,
    clearAllNames,
  }: Props = $props();
</script>

<div class="bg-white/10 backdrop-blur-sm rounded-2xl p-4 md:p-6">
  <div class="flex items-center justify-between mb-4">
    <h2 class="text-xl md:text-2xl font-bold text-white">
      Participants ({names.length})
    </h2>
    {#if names.length > 0}
      <button
        onclick={clearAllNames}
        disabled={isSpinning}
        class="bg-red-500 hover:bg-red-600 text-white px-3 md:px-4 py-1.5 md:py-2 rounded-lg font-semibold text-xs md:text-sm disabled:opacity-50 disabled:cursor-not-allowed transition-colors cursor-pointer"
      >
        🗑️ Clear
      </button>
    {/if}
  </div>

  <!-- Add name input -->
  <div class="mb-4 md:mb-6">
    <div class="flex gap-2">
      <input
        type="text"
        bind:value={currentName}
        onkeypress={handleKeyPress}
        placeholder="Enter a name..."
        maxlength="20"
        class="flex-1 px-3 md:px-4 py-2 md:py-3 rounded-lg bg-white/20 text-white placeholder-gray-300 border-2 border-white/30 focus:border-pink-400 focus:outline-none text-sm md:text-base"
      />
      <button
        onclick={addName}
        disabled={!currentName.trim() || names.length >= 20}
        class="bg-green-500 hover:bg-green-600 text-white px-4 md:px-6 py-2 md:py-3 rounded-lg font-semibold text-sm md:text-base disabled:opacity-50 disabled:cursor-not-allowed transition-colors cursor-pointer"
      >
        Add
      </button>
    </div>
    <p class="text-gray-300 text-xs md:text-sm mt-2">Maximum 20 names</p>
  </div>

  <!-- Names list -->
  <div class="space-y-2 max-h-75 md:max-h-100 overflow-y-auto custom-scrollbar">
    {#if names.length === 0}
      <div class="text-center py-8 md:py-12 text-gray-400">
        <div class="text-4xl md:text-6xl mb-3 md:mb-4">📝</div>
        <p class="text-sm md:text-base">No names added yet</p>
      </div>
    {:else}
      {#each names as name, i}
        <div
          class="flex items-center justify-between bg-white/20 backdrop-blur-sm rounded-lg p-3 md:p-4 hover:bg-white/30 transition-colors"
        >
          <div class="flex items-center gap-2 md:gap-3 overflow-hidden">
            <div
              class="w-7 h-7 md:w-8 md:h-8 rounded-full flex items-center justify-center text-white font-bold text-sm md:text-base shrink-0"
            >
              {i + 1}.
            </div>
            <span class="text-white font-medium text-sm md:text-lg truncate"
              >{name}</span
            >
          </div>
          <button
            onclick={() => removeName(i)}
            disabled={isSpinning}
            class="text-black hover:text-gray-300 font-bold text-xl md:text-2xl disabled:opacity-50 disabled:cursor-not-allowed shrink-0 cursor-pointer"
            title="Remove"
          >
            X
          </button>
        </div>
      {/each}
    {/if}
  </div>
</div>

<style>
  .custom-scrollbar::-webkit-scrollbar {
    width: 8px;
  }

  .custom-scrollbar::-webkit-scrollbar-track {
    background: rgba(255, 255, 255, 0.1);
    border-radius: 10px;
  }

  .custom-scrollbar::-webkit-scrollbar-thumb {
    background: rgba(255, 255, 255, 0.3);
    border-radius: 10px;
  }

  .custom-scrollbar::-webkit-scrollbar-thumb:hover {
    background: rgba(255, 255, 255, 0.5);
  }
</style>

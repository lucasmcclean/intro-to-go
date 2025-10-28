<script lang="ts">
  import { writable } from 'svelte/store';
  import { BACKEND_URL } from './consts';

  let duplicateInput = '';
  const duplicateResult = writable<number[] | null>(null);
  const duplicateError = writable<string | null>(null);

  async function sendDuplicate() {
    duplicateError.set(null);
    duplicateResult.set(null);
    try {
      let entries = duplicateInput
        .split(',')
        .map(e => parseInt(e.trim()))
        .filter(e => !isNaN(e));
      let res = await fetch(`${BACKEND_URL}/duplicate`, {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify({ entries })
      });
      if (!res.ok) throw new Error(await res.text());
      let data = await res.json();
      duplicateResult.set(data.entries);
    } catch (err: any) {
      duplicateError.set(err.message);
    }
  }

  let left = '';
  let right = '';
  let how: 'left-first' | 'right-first' | 'weave' = 'left-first';
  const mixResult = writable<string | null>(null);
  const mixError = writable<string | null>(null);

  async function sendMix() {
    mixError.set(null);
    mixResult.set(null);
    try {
      let res = await fetch(`${BACKEND_URL}/mix`, {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify({ left, right, how })
      });
      if (!res.ok) throw new Error(await res.text());
      let data = await res.json();
      mixResult.set(data.result);
    } catch (err: any) {
      mixError.set(err.message);
    }
  }

  let workers = 3;
  let jobs = 10;
  const concurrencyLog = writable<string[] | null>(null);
  const concurrencyError = writable<string | null>(null);

  async function runConcurrency() {
    concurrencyError.set(null);
    concurrencyLog.set(null);
    try {
      let res = await fetch(`${BACKEND_URL}/jobs`, {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify({ workers, jobs })
      });
      if (!res.ok) throw new Error(await res.text());
      let data = await res.json();
      concurrencyLog.set(data);
    } catch (err: any) {
      concurrencyError.set(err.message);
    }
  }
</script>

<main class="p-8 space-y-12 max-w-xl mx-auto">
  <h1 class="text-3xl font-bold mb-6">Intro to Go API Tester</h1>

  <section class="space-y-2">
    <h2 class="text-xl font-semibold">Duplicate Entries</h2>
    <input
      class="border p-2 w-full"
      placeholder="Enter numbers comma-separated, e.g., 1,2,3"
      bind:value={duplicateInput}
    />
    <button onclick={sendDuplicate}>
      Send
    </button>
    {#if $duplicateResult}
      <pre>{$duplicateResult}</pre>
    {:else if $duplicateError}
      <p class="text-red-500">{$duplicateError}</p>
    {/if}
  </section>

  <!-- Mix -->
  <section class="space-y-2">
    <h2 class="text-xl font-semibold">Mix Strings</h2>
    <input class="border p-2 w-full" placeholder="Left" bind:value={left} />
    <input class="border p-2 w-full" placeholder="Right" bind:value={right} />
    <select class="border p-2 w-full" bind:value={how}>
      <option value="left-first">Left First</option>
      <option value="right-first">Right First</option>
      <option value="weave">Weave</option>
    </select>
    <button onclick={sendMix}>
      Mix
    </button>
    {#if $mixResult}
      <pre>{$mixResult}</pre>
    {:else if $mixError}
      <p class="text-red-500">{$mixError}</p>
    {/if}
  </section>

  <!-- Concurrency Demo -->
  <section class="space-y-2">
    <h2 class="text-xl font-semibold">Concurrency Demo</h2>
    <input type="number" min="1" class="border p-2 w-full" bind:value={workers} placeholder="Workers" />
    <input type="number" min="1" class="border p-2 w-full" bind:value={jobs} placeholder="Jobs" />
    <button onclick={runConcurrency}>
      Run
    </button>
    {#if $concurrencyLog}
      <pre>{$concurrencyLog.join('\n')}</pre>
    {:else if $concurrencyError}
      <p class="text-red-500">{$concurrencyError}</p>
    {/if}
  </section>
</main>

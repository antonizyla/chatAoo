<script lang="ts">
	export let title: string;
	export let content: string;
	export let open = false;

	function toggle_closed(): void {
		open = !open;
	}

	// make it look nice
	import Chevron_down from './chevron-down.svelte';
	import Chevron_up from './chevron-up.svelte';
	import { fade } from 'svelte/transition';
</script>

<div class="">
	<button
		on:click={() => toggle_closed()}
		class={['flex w-full flex-row justify-between p-4'].join(' ')}
	>
		<div class="">{title}</div>
		<div class="h-[32px] w-[32px]">
			{#if open}
				<div class="" in:fade|global={{ duration: 500 }}>
					<Chevron_up />
				</div>
			{:else}
				<div in:fade|global={{ duration: 500 }}>
					<Chevron_down />
				</div>
			{/if}
		</div>
	</button>
	{#if open}
		<div class="px-6 pb-6">{content}</div>
	{/if}
</div>

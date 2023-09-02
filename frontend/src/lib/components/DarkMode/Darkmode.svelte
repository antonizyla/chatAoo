<script lang="ts">
	import { onMount } from 'svelte';
	import { writable, type Writable } from 'svelte/store';

	let darkMode: Writable<boolean> = writable(false);

	import MoonIcon from './MoonIcon.svelte';
	import SunIcon from './SunIcon.svelte';

	// use local storage to persist dark mode into a store
	onMount(() => {
		if (
			localStorage.theme === 'dark' ||
			(!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)
		) {
			setDarkMode();
		} else {
			setLightMode();
		}
	});

	function setDarkMode() {
		$darkMode = true;
		localStorage.setItem('theme', 'dark');
		document.documentElement.classList.add('dark');
	}

	function setLightMode() {
		$darkMode = false;
		localStorage.setItem('theme', 'light');
		document.documentElement.classList.remove('dark');
	}
</script>

<button
	class="m-2 p-2"
	on:click={() => {
		$darkMode ? setLightMode() : setDarkMode();
	}}
>
	{#if $darkMode}
		<div class="flex flex-row gap-3">
			<SunIcon />
			<div class="ml-4">Use Light Mode</div>
		</div>
	{:else}
		<div class="flex flex-row gap-3">
			<MoonIcon />
			<div class="ml-4">Use Dark Mode</div>
		</div>
	{/if}
</button>

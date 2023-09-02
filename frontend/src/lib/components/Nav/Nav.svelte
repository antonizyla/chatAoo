<script lang="ts">
	import Button from '$lib/components/Button/Button.svelte';

	import HamburgerMenu from './HamburgerMenu.svelte';
	import MinimiseHamburgerMenu from './MinimiseHamburger.svelte';

	export let mainLinks: { href: string; label: string }[] = [];
	export let ctaLinks: { href: string; label: string; primary: boolean }[] = [];
	export let title: string;
	export let classes: string = '';
	export let sticky: boolean = false;

	let open: boolean = false;

	function toggleNav() {
		open = !open;
	}
</script>

<nav
	class={'p-4 text-text' + classes}
	class:bg-background={open}
	class:h-screen={open}
	class:sticky
>
	<div class="mx-auto max-w-6xl px-4">
		<div class="flex justify-between">
			<div class="flex space-x-7">
				<div>
					<!-- Website Logo -->
					<a href="/" class="flex items-center px-2 py-4">
						<h2>{title}</h2>
					</a>
				</div>
			</div>
			<!-- Primary Navbar items -->
			<div class="hidden items-center space-x-1 md:flex">
				{#each mainLinks as item}
					<a href={item.href} class="px-2 py-4 font-semibold">{item.label}</a>
				{/each}
			</div>
			<!-- Secondary Navbar items (cta etc..)) -->
			<div class="hidden items-center space-x-3 md:flex">
				{#each ctaLinks as ctaItem}
					<Button primary={ctaItem.primary} size="small">{ctaItem.label}</Button>
				{/each}
			</div>
			<!-- Mobile menu button -->
			<div class="flex items-center md:hidden">
				<button on:click={toggleNav} class="mobile-menu-button outline-none">
					{#if !open}
						<HamburgerMenu />
					{:else}
						<MinimiseHamburgerMenu />
					{/if}
				</button>
			</div>
		</div>
	</div>
	<!-- mobile menu -->
	<div class:hidden={!open} class="absolute top-[35%] h-[100vh] px-4">
		<ul class="">
			{#each mainLinks as item}
				<li>
					<div class="block p-2">
						<a href={item.href}>{item.label}</a>
					</div>
				</li>
			{/each}
			{#each ctaLinks as ctaItem}
				<li>
					<div class="block py-2">
						<Button primary={ctaItem.primary} size="small">{ctaItem.label}</Button>
					</div>
				</li>
			{/each}
		</ul>
	</div>
</nav>

<style lang="postcss">
	.sticky {
		position: sticky;
		top: 0;
		z-index: 100;
		@apply bg-background pb-4;
	}
</style>

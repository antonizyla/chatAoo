<script lang="ts">
	import type { ActionData, PageData } from './$types';
	export let data: PageData;
</script>

<!-- svelte-ignore non-top-level-reactive-declaration -->
{#if data.exists}
	Chat exists with id {data.chat.id} and name {data.chat.name} with description {data.chat
		.description}
	<div>Chat goes here</div>
{:else}
	<h1>404</h1>
	<p>Chat does not exist</p>
{/if}

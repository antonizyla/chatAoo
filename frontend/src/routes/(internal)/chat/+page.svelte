<script lang="ts">
	import { onMount } from 'svelte';
	import type { ActionData } from './$types';
	import { enhance } from '$app/forms';
	import Button from '$lib/components/Button/Button.svelte';
	export let form: ActionData;

	let existingChats: any[] = [];
	let userID: string | null;
	onMount(async () => {
		userID = localStorage.getItem('userID');
		if (!userID) {
			window.location.href = '/';
			return;
		}
		// get chats that the user is inside of
		existingChats = await fetch(`http://localhost:8081/chats/linked/${userID}`).then((res) =>
			res.json()
		);
	});
</script>

<div class="flex flex-row gap-2 p-2">
	<a href="/"><Button size="small">Navigate to Home</Button></a>
	<a href="/account"><Button size="small">Edit Your Account Details</Button></a>
	<a href="/chat/create"><Button size="small">Create a New Chat</Button></a>
</div>

{#if existingChats.length > 0}
	<h2>Chats you're already in</h2>
	{#each existingChats as chat}
		<div class="">
			<p>Name: {chat.name}, Description: {chat.description}</p>
			<button><a href={`/chat/${chat.id}`}>Launch Chat</a></button>
			<button
				on:click={() => {
					existingChats = existingChats.filter((c) => c.id !== chat.id);
					// api call to remove from database
				}}>Leave Chat</button
			>
			<button
				on:click={() => {
					// write to clipboard
					navigator.clipboard.writeText(chat.id);
				}}>Share Chat ID</button
			>
		</div>
	{/each}
{/if}

<style lang="postcss">
	input {
		@apply border border-blue-400 p-1;
	}
</style>

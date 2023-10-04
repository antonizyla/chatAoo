<script lang="ts">
	import { onMount } from 'svelte';
	import type { ActionData } from './$types';
	import Button from '$lib/components/Button/Button.svelte';

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
	<div class="flex flex-col gap-2 p-2">
		{#each existingChats as chat}
			<div class="p-2 bg-gray-100 flex flex-row gap-0.5 justify-between">
				<div class="flex flex-col gap-1.5">
					<div>Name: {chat.name}</div>
					<div>Description: {chat.description}</div>
				</div>
				<div class="text-right self-end">
					<a href={`/chat/${chat.id}`}><Button size="small">Open Chat</Button></a>
					<Button
						size="small"
						on:click={() => {
							existingChats = existingChats.filter((c) => c.id !== chat.id);
							// api call to remove from database
						}}>Leave Chat</Button
					>
					<Button
						size="small"
						on:click={() => {
							// write to clipboard
							navigator.clipboard.writeText(chat.id);
						}}>Share Chat ID</Button
					>
				</div>
			</div>
		{/each}
	</div>
{/if}

<style lang="postcss">
	input {
		@apply border border-blue-400 p-1;
	}
</style>

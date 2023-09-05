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

<h1>Welcome to Chat App</h1>
{#if !form?.success}
	<form class="m-4" action="?/createChat" method="POST" use:enhance>
		<h2>Create a New Chat</h2>
		<input type="text" name="name" placeholder="Chat Name" />
		<input type="text" name="description" placeholder="Chat Description" />
		<input type="text" name="userID" bind:value={userID} hidden />
		<button type="submit"> Create Chat </button>
	</form>
	<form class="m-4" action="?/joinChat" method="POST" use:enhance>
		<h2>Join an existing Chat</h2>
		<input type="text" name="uuid" placeholder="unique chat id" />
		<input type="text" hidden bind:value={userID} name="userID" />
		<button type="submit">Join Chat</button>
	</form>
{:else}
	Successfully joined Chat with name {form?.name} and url: /chat/{form.chat_id}
	<button><a href={`/chat/${form.chat_id}`}>Launch Chat</a></button>
{/if}

<div class="flex flex-row gap-2 p-2">
	<a href="/"><Button size="small">Navigate to Home</Button></a>
	<a href="/account"><Button size="small">Edit Your Account Details</Button></a>
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
		@apply border border-blue-400 p-0.5;
	}

	button {
		@apply border border-blue-600 p-0.5;
	}
</style>

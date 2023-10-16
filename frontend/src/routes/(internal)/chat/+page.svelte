<script lang="ts">
	import { onMount } from 'svelte';
	import Button from '$lib/svelte-components/components/Button/Button.svelte';

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

<div class="flex flex-row gap-2 p-4 mx-auto w-fit">
	<a href="/"><Button size="small">Navigate to Home</Button></a>
	<a href="/account"><Button size="small">Edit Your Account Details</Button></a>
	<a href="/chat/create"><Button size="small">Create a New Chat</Button></a>
</div>

{#if existingChats.length > 0}
	<div class="flex flex-col gap-2 p-2 mx-20">
		{#each existingChats as chat}
			<div
				class="p-2 border rounded-md shadow drop-shadow-sm shadow-primary/50 border-solid flex flex-row gap-0.5 justify-between"
			>
				<div class="flex flex-col gap-1.5">
					<div class="text-lg inline">{chat.name}</div>
					<div class="opacity-70">{chat.description}</div>
				</div>
				<div class="text-right self-end">
					<a href={`/chat/${chat.id}`}><Button primary size="small">Open Chat</Button></a>
					<Button
						size="small"
						on:click={async () => {
							let unlinking = await fetch(`http://localhost:8081/chats/link/${chat.id}/${userID}`, {
								method: 'DELETE'
							}).then((res) => res.status);
							console.log(unlinking);
							if (unlinking == 200) {
								existingChats = existingChats.filter((c) => c.id !== chat.id);
							}
							// api call to remove from database
						}}>Leave Chat</Button
					>
					<Button
						size="small"
						on:click={() => {
							// write to clipboard
							navigator.clipboard.writeText(chat.id);
							alert('Chat ID Copied to Clipboard');
						}}>Share Chat</Button
					>
				</div>
			</div>
		{/each}
	</div>
{:else}
	<div class="flex-col flex mx-auto w-fit mx-auto pt-20 gap-2 text-center">
		<h2>You're not currently in any chats, Join or create one</h2>
		<a href="/chat/create"><Button primary>Join or Create a Chat Now</Button></a>
	</div>
{/if}

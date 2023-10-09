<script lang="ts">
	import { onMount } from 'svelte';
	import type { ActionData } from './$types';
	import { enhance } from '$app/forms';
	import Button from '$lib/components/Button/Button.svelte';
	export let form: ActionData;

	let userID: string;
	onMount(async () => {
		const storedUser = localStorage.getItem('userID');
		if (!storedUser) {
			location.href = '/';
		} else {
			userID = storedUser;
		}
	});

	$: if (form?.success) {
		window.alert('Successfully joined the chat');
		location.href = '/chat';
	}
</script>

<a class="p-2 block" href="/chat">
	<Button primary size="small">Home</Button>
</a>

<div class="mx-auto w-fit pt-40">
	<h2 class="text-center text-xl">Join or Create a Chat</h2>
	<form class="m-4 w-fit mx-auto" action="?/createChat" method="POST" use:enhance>
		<div class="m-2 border border-primary border-solid p-2 rounded-md">
			<h2 class="pl-1 pt-1 pb-2">Create a New Chat</h2>
			<div class="flex flex-row justify-between gap-2">
				<div class="flex flex-col gap-2">
					<input
						type="text"
						class="p-1.5 bg-secondary rounded-md shadow-primary drop-shadow-sm"
						name="name"
						placeholder="Chat Name"
					/>
					<input
						type="text"
						class="p-1.5 bg-secondary rounded-md shadow-primary drop-shadow-sm"
						name="description"
						placeholder="Chat Description"
					/>
				</div>
				<input type="text" name="userID" bind:value={userID} hidden />
				<button type="submit">
					<Button primary size="small">Create Chat</Button>
				</button>
			</div>
		</div>
	</form>
	<form action="?/joinChat" method="POST" use:enhance class="m-4 w-fit mx-auto">
		<div class="m-2 flex flex-col border border-primary border-solid p-2 rounded-md">
			<h2 class="pl-1 pt-1 pb-2">Join an existing Chat</h2>
			<div class="flex flex-row justify-around items-stretch gap-2">
				<input
					type="text"
					class="flex-grow-3 p-1.5 bg-secondary rounded-md shadow-primary drop-shadow-sm"
					name="uuid"
					placeholder="unique chat id"
				/>
				<input type="text" hidden bind:value={userID} name="userID" />
				<button type="submit" class="flex-grow-1">
					<Button size="small" primary>Join a Chat</Button>
				</button>
			</div>
		</div>
	</form>
</div>

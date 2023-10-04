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

<form class="m-4" action="?/createChat" method="POST" use:enhance>
	<h2>Create a New Chat</h2>
	<input type="text" name="name" placeholder="Chat Name" />
	<input type="text" name="description" placeholder="Chat Description" />
	<input type="text" name="userID" bind:value={userID} hidden />
	<button type="submit">
		<Button size="small">Create Chat</Button>
	</button>
</form>
<form class="m-4" action="?/joinChat" method="POST" use:enhance>
	<h2>Join an existing Chat</h2>
	<input type="text" name="uuid" placeholder="unique chat id" />
	<input type="text" hidden bind:value={userID} name="userID" />
	<button type="submit">
		<Button size="small">Join Chat</Button>
	</button>
</form>

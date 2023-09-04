<script lang="ts">
	import Button from '$lib/components/Button/Button.svelte';
	import { onMount } from 'svelte';

	let currentUser: string = '';
	let currentUserName: string = '';
	let displayName: string = '';
	let accountExists: boolean = true;
	onMount(() => {
		if (!localStorage.getItem('userID')) {
			accountExists = false;
		}
		currentUser = localStorage.getItem('userID') || '';
		currentUserName = localStorage.getItem('userName') || '';
		displayName = currentUserName;
	});

	async function changeUsername() {
		const update = await fetch(`http://localhost:8081/users/${currentUser}`, {
			method: 'PATCH',
			body: JSON.stringify({ name: currentUserName })
		});
		console.log(update);
		if (update.status === 200) {
			alert('Username Updated');
			localStorage.setItem('userName', currentUserName);
			displayName = currentUserName;
		}
	}
</script>

{#if accountExists}
	<p>Your Account Exists, Edit Your details Below</p>
	<div>
		You are currently signed in as <div class="bg-red-50 inline">{displayName}</div>
		with internal user id
		<div class="bg-red-50 inline">{currentUser}</div>
		<div class="">
			Current UserName:
			<input
				class="bg-red-50 p-2"
				type="text"
				name="userName"
				id="userName"
				bind:value={currentUserName}
			/>
			<button on:click={changeUsername}>Change Username</button>
		</div>
	</div>
{:else}
	<p>You do not have an account stored in your browser</p>
	<a href="/"><Button>Create an Account</Button></a>
{/if}

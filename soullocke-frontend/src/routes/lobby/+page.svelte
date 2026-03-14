<script lang="ts">
	import Card from '../../components/Card.svelte';
	import PrimaryButton from '../../components/buttons/PrimaryButton.svelte';
	import TextInput from '../../components/TextInput.svelte';
	import { createLobbyQuery } from '../../api/lobby.ts';

	let lobbyId = $state('');
	const lobbyQuery = createLobbyQuery(() => lobbyId);

	let lobbyName = $state('Test Lobby');

	const handleBtnClick = (): void => {
		console.log(lobbyName);
	};


</script>

<h1 class="text-center text-3xl">
	Lobby Management
</h1>



<div class="grid grid-cols-3 mt-10 gap-4">
	<div class="col-start-2">
		<Card>
			<header slot="header">
				<h2 class="h4">Create Lobby</h2>
			</header>
			<section slot="content">
				<TextInput label="Lobby Name" bind:value={lobbyName} />
			</section>
			<section slot="footer">
				<PrimaryButton variant="filled" onClick={handleBtnClick}>Create Lobby</PrimaryButton>
			</section>
		</Card>
	</div>
	<div class="col-start-2">
		<Card>
			<header slot="header">
				<h2 class="h4">Get Lobby</h2>
			</header>
			<section slot="content">
				<TextInput label="Lobby Name" bind:value={lobbyId} />
				{#if lobbyQuery.isSuccess}
					{lobbyQuery.data.name}
				{/if}
				{#if lobbyQuery.isError}
					<p class="text-red-500">Error fetching lobby</p>
				{/if}
			</section>
			<section slot="footer">
				<PrimaryButton variant="filled" onClick={handleBtnClick}>Create Lobby</PrimaryButton>
			</section>
		</Card>
	</div>
</div>

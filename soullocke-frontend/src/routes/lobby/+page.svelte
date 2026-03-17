<script lang="ts">
	import Card from '$components/Card.svelte';
	import PrimaryButton from '$components/buttons/PrimaryButton.svelte';
	import TextInput from '$components/textinput/TextInput.svelte';
	import { getLobbyCreationMutation } from '$api/lobby';
	import Combobox from '$components/Combobox.svelte';

	let lobbyName = $state('');
	let lobbyPassword = $state('')
	let gameEdition = $state<string>();
	let lobbyNameValid = $state(false);
	let lobbyPasswordValid = $state(false);

	const createLobbyMutation = getLobbyCreationMutation();

	const handleCreateLobbyClick = (): void => {
		if (!gameEdition) {
			return;
		}
		createLobbyMutation.mutate({
			body: {
				name: lobbyName.trim(),
				password: lobbyPassword,
				gameEditionId: gameEdition,
			}
		});
	};

</script>

<h1 class="text-center text-3xl">
	Lobby Management
</h1>

<div class="mx-auto flex max-w-1/2 flex-col gap-4 mt-10">
	<Card>
		{#snippet header()}
			<h2 class="h4">Create Lobby</h2>
		{/snippet}
		{#snippet content()}
			<div class="flex flex-col gap-4">
				<TextInput
					label="Lobby Name"
					disabled={createLobbyMutation.isPending}
					minLength={10}
					maxLength={255}
					bind:value={lobbyName}
					bind:valid={lobbyNameValid}
				/>
				<TextInput
					label="Password"
					type="password"
					disabled={createLobbyMutation.isPending}
					minLength={8}
					bind:value={lobbyPassword}
					bind:valid={lobbyPasswordValid}
				/>
				<Combobox
					items={[
						{ label: 'FireRed', value: 'firered', group: 'Generation 3 - Remake' },
						{ label: 'LeafGreen', value: 'leafgreen', group: 'Generation 3 - Remake' },
						{ label: 'Emerald', value: 'emerald', group: 'Generation 3' },
						{ label: 'Ruby', value: 'ruby', group: 'Generation 3' },
						{ label: 'Sapphire', value: 'sapphire', group: 'Generation 3' },
					]}
					label="Game Edition"
					bind:value={gameEdition}
				/>
			</div>
		{/snippet}
		{#snippet footer()}
			<PrimaryButton
				variant="filled"
				disabled={createLobbyMutation.isPending || !lobbyNameValid || !lobbyPasswordValid || !gameEdition}
				onClick={handleCreateLobbyClick}
			>
				Create Lobby
			</PrimaryButton>
		{/snippet}
	</Card>
	{#if createLobbyMutation.isSuccess && createLobbyMutation.data}
		<div class="col-start-2">
			<Card>
				{#snippet header()}
					<h2 class="h4">Lobby Created</h2>
				{/snippet}
				{#snippet content()}
					<p>Lobby ID: {createLobbyMutation.data.lobbyId}</p>
				{/snippet}
			</Card>
		</div>
	{/if}
	{#if createLobbyMutation.isError}
		<div class="col-start-2">
			<Card>
				{#snippet header()}
					<h2 class="h4">Error Creating Lobby</h2>
				{/snippet}
				{#snippet content()}
					<p class="text-error-500">Error Error Error :(</p>
				{/snippet}
			</Card>
		</div>
	{/if}
</div>

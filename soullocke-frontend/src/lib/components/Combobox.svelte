<script lang="ts" generics="T extends unknown">
	import { Combobox, Portal, type ComboboxRootProps, useListCollection } from '@skeletonlabs/skeleton-svelte';

	type ComboboxItem = {
		label: string;
		value: T;
		group?: string;
	};

	type Props = {
		label: string;
		items: ComboboxItem[];
		value?: string;
		disabled?: boolean;
	};

	let {
		label,
		items,
		disabled = false,
		value = $bindable()
	} : Props = $props();
	const comboboxValue = $derived(value ? [value] : []);

	let filterText = $state<string>()

	const collection = $derived(
		useListCollection({
			items: items.filter((item) => item.label.toLowerCase().includes(filterText?.toLowerCase() ?? '')),
			groupBy: (item) => item.group ?? '',
		}),
	);

	const onOpenChange = () => {
		filterText = '';
	};

	const onInputValueChange: ComboboxRootProps['onInputValueChange'] = (event) => {
		filterText = event.inputValue;
	};
</script>

<Combobox
	class="max-w"
	placeholder="Search..."
	{collection}
	{onOpenChange}
	{onInputValueChange}
	{disabled}
	value={comboboxValue}
	multiple={false}
	onValueChange={(e) => { value = e.value.length === 1 ? e.value[0] : undefined }}
>
	<Combobox.Label>{label}</Combobox.Label>
	<Combobox.Control>
		<Combobox.Input class="p-3 rounded-md ring-1 hover:ring-surface-400"/>
		<Combobox.Trigger />
	</Combobox.Control>
	<Combobox.ClearTrigger>Clear</Combobox.ClearTrigger>
	<Portal>
		<Combobox.Positioner>
			<Combobox.Content>
				{#each collection.group() as [type, items] (type)}
					<Combobox.ItemGroup>
						<Combobox.ItemGroupLabel>{type}</Combobox.ItemGroupLabel>
						{#each items as item (item.value)}
							<Combobox.Item {item} class="p-2">
								<Combobox.ItemText>{item.label}</Combobox.ItemText>
								<Combobox.ItemIndicator />
							</Combobox.Item>
						{/each}
					</Combobox.ItemGroup>
				{/each}
			</Combobox.Content>
		</Combobox.Positioner>
	</Portal>
</Combobox>

<script lang="ts" generics="T extends unknown">
	import { Combobox, Portal, type ComboboxRootProps, useListCollection } from '@skeletonlabs/skeleton-svelte';

	type Props = {
		label: string;
		items: { label: string; value: T, group?: string }[];
		value?: string;
	};

	let {
		label,
		items,
		value = $bindable()
	} : Props = $props();
	const comboboxValue = $derived(value ? [value] : []);

	let filteredItems = $state(items);

	const collection = $derived(
		useListCollection({
			items: filteredItems,
			groupBy: (item) => item.group ?? '',
		}),
	);

	const onOpenChange = () => {
		filteredItems = items;
	};

	const onInputValueChange: ComboboxRootProps['onInputValueChange'] = (event) => {
		const filtered = items.filter((item) => item.label.toLowerCase().includes(event.inputValue.toLowerCase()));
		if (filtered.length > 0) {
			filteredItems = filtered;
		} else {
			filteredItems = items;
		}
	};
</script>

<Combobox
	class="max-w"
	placeholder="Search..."
	{collection}
	{onOpenChange}
	{onInputValueChange}
	value={comboboxValue}
	multiple={false}
	onValueChange={(e) => { value = e.value.length === 1 ? e.value[0] : undefined }}
>
	<Combobox.Label>{label}</Combobox.Label>
	<Combobox.Control>
		<Combobox.Input />
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
							<Combobox.Item {item}>
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

import type { Meta, StoryObj } from '@storybook/svelte';

import Cards from './Cards.svelte';

const meta = {
	title: 'Cards',
	component: Cards,
	tags: ['autodoc'],
	argTypes: {}
} satisfies Meta<Cards>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Primary: Story = {
	args: {
		contents: [
			{
				title: 'Card Title Goes Here...',
				content:
					'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla vitae elit libero, a pharetra augue.'
			},
			{
				title: 'Card Title Goes Here...',
				content:
					'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla vitae elit libero, a pharetra augue.'
			},
			{
				title: 'Card Title Goes Here...',
				content:
					'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla vitae elit libero, a pharetra augue.'
			}
		]
	}
};

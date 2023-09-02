import type { Meta, StoryObj } from '@storybook/svelte';

import Accordion from './Accordion.svelte';

const meta = {
	title: 'Accordion',
	component: Accordion,
	tags: ['autodocs'],
	argTypes: {}
} satisfies Meta<Accordion>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Primary: Story = {
	args: {
		accordionItems: [
			{
				id: '1',
				title: 'Accordion Item 1',
				content:
					'Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.',
				open: true
			},
			{
				id: '2',
				title: 'Accordion Item 2',
				content:
					'Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.',
				open: false
			},
			{
				id: '3',
				title: 'Accordion Item 3',
				content:
					'Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.',
				open: false
			},
			{
				id: '4',
				title: 'Accordion Item 4',
				content:
					'Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.',
				open: false
			}
		]
	}
};

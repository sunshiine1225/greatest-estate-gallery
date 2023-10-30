import {defineType} from 'sanity'

export default {
	name: 'property',
	type: 'document',
	title: 'Property',
	fields: [
		defineType({
			title: 'Property Name',
			name: 'property_name',
			type: 'string',
		}),
		defineType({
			title: 'Images',
			name: 'images',
			type: 'array',
			of: [
				{
					type: 'image',
				},
			],
		}),
		defineType({
			name: 'type',
			title: 'Property Type',
			type: 'string',
			options: {
				list: [
                    '1 BHK',
                    '1-2 BHK',
                    '2 BHK',
                    '2-3 BHK',
                    '3 BHK',
                    '1-3 BHK',
                    '4 BHK',
                    '2-4 BHK',
                    '3-4 BHK',
                ],
				layout: 'dropdown',
			},
		}),
        defineType({
            type: 'object',
            name: 'price_range',
            title: 'Price Range',
            fields:  [
                { type: 'number', name: 'lower_limit', title: 'Lower Limit'},
                { type: 'number', name: 'upper_limit', title: 'Upper Limit'},
            ]
        }),
		defineType({
			name: 'full_desc',
			title: 'Description',
			type: 'string',
		}),
	],
}

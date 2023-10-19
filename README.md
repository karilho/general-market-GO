# general-market-GO
This is a general purpose use market Backend API.

## How it was planned?
<a href="https://ibb.co/SVpwnmX"><img src="https://i.ibb.co/BTR2Z6V/template.png" alt="template" border="0"></a>

First, we have 2 types of users: Buyers and Sellers.

Buyers can buy products directly or from Sellers (Not implemented yet), and this generates a buyer_order.

Buyer order table references the purchased products, and this stable stores the quantity of each product and price 
of them at the moment of the purchase.

And we have too product table, that stores the products that are available to be purchased.

 


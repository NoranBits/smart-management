= Functional Requirements for Smart Management App

== 1. Overview
This document defines the functional requirements for a restaurant management system designed for both customers and staff. The system will allow customers to browse menus, place orders, track order status, and make payments, while restaurant staff (waiters, chefs, managers) will manage orders, track inventory, update order statuses, and generate reports.

== 2. User Roles and Access Control

=== 2.1 Customers
Customers interact with the system through the mobile application (Android) to perform the following actions:

Register and log in (optional for guests).

Browse the restaurant’s menu and view item details.

Select items, customize orders, and place orders.

Assign their order to a table via visually enchanted manual selection.  (QR code scanning considered later)

Track order status in real-time.

Make online payments via credit/debit cards, digital wallets, or opt for cash payments.

Receive push notifications about order progress.

View order history and request receipts.

=== 2.2 Staff (Waiters, Chefs, and Managers) The staff interacts with the system through the desktop application (Windows/Linux) or mobile/tablet application(android) to manage the restaurant operations:

Waiters:

Accept and manage customer orders.

Assign or update table numbers for on-site orders.

Notify the kitchen about new orders.

Change order status (e.g., from Pending to In Progress to Served).

Process payments and generate customer invoices.

View order history and customer preferences.

Chefs:

View new incoming orders and prioritize food preparation.

Mark orders as "Ready for Pickup" when completed.

Send notifications to waiters once an order is prepared.

Managers:

Monitor overall restaurant performance through analytics and reports.

Manage inventory levels and reorder supplies when needed.

Assign roles and permissions to staff members.

View detailed financial reports and daily sales.

== 3. Functional Features

=== 3.1 Authentication and User Management

Customers can register via email or phone number.

Staff accounts are created and assigned roles by an administrator.

JWT-based authentication for secure access.

Password reset functionality for all users.

Two-factor authentication (option is considered later)

=== 3.2 Menu Management

Admins and managers can add, edit, and remove menu items.

Each menu item includes name, description, price, and availability.

Menu items are categorized (e.g., Starters, Main Course, Desserts, Beverages).

Ability to upload images for menu items.

=== 3.3 Ordering System

Customers can add multiple items to their cart before placing an order.

Orders can be for:

Dine-in (assigned to a table via QR code or manual input).

Takeout (customer picks up the order at the counter).

Delivery (if supported by the restaurant).

Customers can select special instructions for food preparation.

Staff can update orders before they reach the preparation stage.

Orders have the following statuses:

Pending

In Progress

Ready for Pickup

Served/Completed

Canceled

=== 3.4 Payment Processing 

Customers can pay via:

Credit/Debit Cards (integrated with PayPal API).

Digital Wallets (Google Pay, Apple Pay, etc.). (option is considered later)

Cash (handled by waiters at checkout).

Secure transaction processing using industry-standard encryption.

Refund and order cancellation policies applied based on restaurant policies.

Receipt generation for each transaction.

=== 3.5 Order Tracking and Notifications

Customers receive real-time updates on their orders.

WebSockets used for real-time status changes. 

Push notifications sent for important updates:

Order confirmed

Order being prepared

Order ready for pickup

Payment successful

=== 3.6 Inventory Management (For Staff Only)

Track stock levels of all menu ingredients.

Automatic updates when items are used in orders.

Alerts when inventory reaches a low threshold.

Generate inventory reports and reorder suggestions.

Prevent ordering items that are out of stock.

=== 3.7 Reports and Analytics (For Managers)

Daily, weekly, and monthly sales reports.

Order trends and customer preferences.

Employee performance tracking.

Stock usage and wastage reports.

Revenue breakdown per order type (dine-in, takeout, delivery).

=== 3.8 Offline Mode for Staff (is considered later).

Orders can be taken offline if the internet is unavailable.

Local SQLite storage for offline orders.

Auto-sync when the internet connection is restored.

Ensures that staff can continue working without disruptions.

== 4. Non-Functional Requirements

=== 4.1 Performance and Scalability

Must handle at least 1000 orders per day.

Go backend optimized with PostgreSQL indexing and connection pooling.

Asynchronous task handling for better API performance.

=== 4.2 Security (is considered later)

JWT authentication and role-based access control.

Secure payment processing with encryption.

Protection against SQL Injection and Cross-Site Scripting (XSS).

Data encryption for sensitive user data.

=== 4.3 Reliability and Availability

System should operate 24/7 without significant downtime.

Auto-recovery for transactions in case of app crashes.

Cloud-based deployment with automatic failover (option is considered later).
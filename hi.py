import pygraphviz as pgv

# Create a new graph
graph = pgv.AGraph(directed=True)

# Define the tables and their relationships
tables = {
    'User': ['UserID', 'Name', 'Email', 'Password', 'CurrentLocation', 'RegistrationDate'],
    'Province': ['ProvinceID', 'Name'],
    'Regency': ['RegencyID', 'Name', 'ProvinceID'],
    'City': ['CityID', 'Name', 'RegencyID'],
    'District': ['DistrictID', 'Name', 'CityID'],
    'MedicalFacility': ['FacilityID', 'Name', 'Location', 'DistrictID', 'ContactNumber', 'OperatingHours'],
    'Weather': ['WeatherID', 'DistrictID', 'Temperature', 'Humidity', 'Rainfall', 'LastUpdated'],
    'Donation': ['DonationID', 'UserID', 'Amount', 'PaymentStatus', 'DonationDate'],
    'Payment': ['PaymentID', 'DonationID', 'PaymentType', 'PaymentDate', 'Amount'],
    'DonationHistory': ['HistoryID', 'UserID', 'DonationID', 'DonationDate']
}

# Add nodes to the graph
for table, columns in tables.items():
    graph.add_node(table, shape='record', label='{}|{}'.format(table, '|'.join(columns)))

# Define the relationships
relationships = [
    ('User', 'Donation'),
    ('Donation', 'Payment'),
    ('User', 'DonationHistory'),
    ('District', 'MedicalFacility'),
    ('District', 'Weather'),
    ('City', 'District'),
    ('Regency', 'City'),
    ('Province', 'Regency')
]

# Add edges to the graph
for relationship in relationships:
    graph.add_edge(relationship[0], relationship[1])

# Set the layout engine
graph.layout(prog='dot')

# Save the graph as an image
graph.draw('erd.png')

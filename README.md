# Projet_Final_API\_\_Clinique_V-t-rinaire_MatiasB_EvanG_MaximeF

Documentation de l'API de Clinique Vétérinaire
Introduction
Cette API permet de gérer les entrées de chats, les visites et les traitements dans une clinique vétérinaire. Elle est construite en utilisant Go et le framework Chi pour le routage.

Endpoints
Chats
Créer un chat
URL: /api/v1/clinique/cat/create-cat
Méthode: POST
Corps de la requête:
{
"cat_name": "Nom du chat",
"cat_age": 3,
"cat_race": "Race du chat",
"cat_weight": 5
}
Réponse:
{
"cat_name": "Nom du chat",
"cat_age": 3,
"cat_race": "Race du chat",
"cat_weight": 5
}

Obtenir un chat par ID
URL: /api/v1/clinique/cat/one-cat/{id}
Méthode: GET
Réponse:
{
"visit_date": "2023-10-01",
"visit_veto": "Nom du vétérinaire",
"visit_motif": "Motif de la visite",
"visit_id_cat": 1
}

Mettre à jour un chat
URL: /api/v1/clinique/cat/update-cat/{id}
Méthode: PUT
Corps de la requête:
{
"cat_name": "Nom du chat",
"cat_age": 4,
"cat_race": "Nouvelle race",
"cat_weight": 6
}
Réponse:
{
"cat_name": "Nom du chat",
"cat_age": 4,
"cat_race": "Nouvelle race",
"cat_weight": 6
}

Supprimer un chat
URL: /api/v1/clinique/cat/delete-cat/{id}
Méthode: DELETE
Réponse:
{
"message": "Oups, nous avons tué votre chat !"
}

Visites
Créer une visite
URL: /api/v1/clinique/visit/visit_create
Méthode: POST
Corps de la requête:
{
"visit_date": "2023-10-01",
"visit_veto": "Nom du vétérinaire",
"visit_motif": "Motif de la visite",
"visit_id_cat": 1
}
Réponse:
{
"visit_date": "2023-10-01",
"visit_veto": "Nom du vétérinaire",
"visit_motif": "Motif de la visite",
"visit_id_cat": 1
}

Obtenir toutes les visites
URL: /api/v1/clinique/visit/all-visits
Méthode: GET
Réponse:
[
{
"visit_date": "2023-10-01",
"visit_veto": "Nom du vétérinaire",
"visit_motif": "Motif de la visite",
"visit_id_cat": 1
},
...
]

Obtenir une visite par ID
URL: /api/v1/clinique/visit/one-visit/{id}
Méthode: GET
Réponse:
{
"visit_date": "2023-10-01",
"visit_veto": "Nom du vétérinaire",
"visit_motif": "Motif de la visite",
"visit_id_cat": 1
}

Mettre à jour une visite
URL: /api/v1/clinique/visit/update-visit/{id}
Méthode: PUT
Corps de la requête:
{
"visit_date": "2023-10-02",
"visit_veto": "Nouveau vétérinaire",
"visit_motif": "Nouveau motif",
"visit_id_cat": 1
}
Réponse:
{
"visit_date": "2023-10-02",
"visit_veto": "Nouveau vétérinaire",
"visit_motif": "Nouveau motif",
"visit_id_cat": 1
}

Supprimer une visite
URL: /api/v1/clinique/visit/delete-visit/{id}
Méthode: DELETE
Réponse:
{
"message": "Oups, nous avons tué votre visite !"
}

Traitements
Créer un traitement
URL: /api/v1/clinique/treatment/treatment_create
Méthode: POST
Corps de la requête:
{
"treatment_medoc": "Nom du médicament",
"treatment_id_visit": 1
}
Réponse:
{
"treatment_medoc": "Nom du médicament",
"treatment_id_visit": 1
}

Obtenir tous les traitements
URL: /api/v1/clinique/treatment/all-treatments
Méthode: GET
Réponse:
[
{
"treatment_medoc": "Nom du médicament",
"treatment_id_visit": 1
},
...
]

Obtenir un traitement par ID
URL: /api/v1/clinique/treatment/one-treatment/{id}
Méthode: GET
Réponse:
{
"treatment_medoc": "Nom du médicament",
"treatment_id_visit": 1
}

Mettre à jour un traitement
URL: /api/v1/clinique/treatment/update-treatment/{id}
Méthode: PUT
Corps de la requête:
{
"treatment_medoc": "Nouveau médicament",
"treatment_id_visit": 1
}
Réponse:
{
"treatment_medoc": "Nouveau médicament",
"treatment_id_visit": 1
}

Supprimer un traitement
URL: /api/v1/clinique/treatment/delete-treatment/{id}
Méthode: DELETE
Réponse:
{
"message": "Oups, nous avons tué votre traitement !"
}

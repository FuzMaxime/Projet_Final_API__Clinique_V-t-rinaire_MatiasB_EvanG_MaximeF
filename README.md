# Documentation de l'API de Clinique Vétérinaire

## Introduction
Cette API permet de gérer les entrées de chats, les visites et les traitements dans une clinique vétérinaire. Elle est construite en utilisant Go et le framework Chi pour le routage.

## Endpoints

## Chats

#### Créer un chat
- **URL**: `/api/v1/clinique/cat/create-cat`
- **Méthode**: `POST`
- **Corps de la requête**:
  ```json
  {
    "cat_name": "Nom du chat",
    "cat_age": 3,
    "cat_race": "Race du chat",
    "cat_weight": 5
  }
- **Réponse**:
  ```json
  {
  "cat_name": "Nom du chat",
  "cat_age": 3,
  "cat_race": "Race du chat",
  "cat_weight": 5
  }

### Obtenir tous les chats
- **URL**: `/api/v1/clinique/cat/all-cats`
- **Méthode**: `GET`
- **Réponse**:
  ```json
  [
    {
        "ID": 1,
        "CreatedAt": "2024-12-02T10:00:32.9010576+01:00",
        "UpdatedAt": "2024-12-02T10:00:32.9010576+01:00",
        "DeletedAt": null,
        "cat_name": "doggo",
        "cat_age": 150,
        "cat_race": "pharaon",
        "cat_weight": 150
    },
    ...
  ]

## Obtenir un chat par ID
- **URL**: `/api/v1/clinique/cat/one-cat/{id}`
- **Méthode**: `GET`
- **Réponse**:
  ```json
  {
    "visit_date": "2023-10-01",
    "visit_veto": "Nom du vétérinaire",
    "visit_motif": "Motif de la visite",
    "visit_id_cat": 1
  }

### Obtenir les visites d'un chat par ID
- **URL**: `/api/v1/clinique/cat/one-cat/{id_cat}/visits`
- **Méthode**: `GET`
- **Réponse**:
  ```json
  [
    {
        "ID": 1,
        "CreatedAt": "2024-12-02T14:43:22.7645512+01:00",
        "UpdatedAt": "2024-12-02T15:25:02.4559814+01:00",
        "DeletedAt": null,
        "visit_date": "02/02/2004",
        "visit_veto": "Evan",
        "visit_motif": "dog",
        "visit_id_cat": 6
    },
    ...
  ]


### Mettre à jour un chat
- **URL**: `/api/v1/clinique/cat/update-cat/{id}`
- **Méthode**: `PUT`
- **Corps de la requête**:
  ```json
  {
    "cat_name": "Nom du chat",
    "cat_age": 4,
    "cat_race": "Nouvelle race",
    "cat_weight": 6
  }
- **Réponse**:
  ```json
  {
  "cat_name": "Nom du chat",
  "cat_age": 4,
  "cat_race": "Nouvelle race",
  "cat_weight": 6
  }
  
## Supprimer un chat
- **URL**: `/api/v1/clinique/cat/delete-cat/{id}`  
- **Méthode**: `DELETE`  
- **Réponse**:
  ```json
  {
    "message": "Oups, nous avons tué votre chat !"
  }

## Visites

### Créer une visite
-**URL**: `/api/v1/clinique/visit/visit_create`  
-**Méthode**: `POST`  
- **Corps de la requête**:
  ```json
  {
    "visit_date": "2023-10-01",
    "visit_veto": "Nom du vétérinaire",
    "visit_motif": "Motif de la visite",
    "visit_id_cat": 1
  }
- **Réponse**:
  ```json
  {
  "visit_date": "2023-10-01",
  "visit_veto": "Nom du vétérinaire",
  "visit_motif": "Motif de la visite",
  "visit_id_cat": 1
  }

## Obtenir toutes les visites
-**URL**: `/api/v1/clinique/visit/all-visits`  
-**Méthode**: `GET`  
- **Réponse**:
  ```json
  [
    {
      "visit_date": "2023-10-01",
      "visit_veto": "Nom du vétérinaire",
      "visit_motif": "Motif de la visite",
      "visit_id_cat": 1
    },
    ...
  ]

## Obtenir une visite par ID
-**URL**: `/api/v1/clinique/visit/one-visit/{id}`  
-**Méthode**: `GET`  
- **Réponse**:
  ```json
  {
    "visit_date": "2023-10-01",
    "visit_veto": "Nom du vétérinaire",
    "visit_motif": "Motif de la visite",
    "visit_id_cat": 1
  }

### Obtenir les traitements d'une visite par ID
- **URL**: `/api/v1/clinique/visit/one-visit/{id_visit}/treatments`
- **Méthode**: `GET`
- **Réponse**:
  ```json
  [
    {
        "ID": 2,
        "CreatedAt": "2024-12-02T15:22:50.8688315+01:00",
        "UpdatedAt": "2024-12-02T15:22:50.8688315+01:00",
        "DeletedAt": null,
        "treatment_medoc": "doliprane",
        "treatment_id_visit": 1
    }
    ...
  ]

## Mettre à jour une visite
-**URL**: `/api/v1/clinique/visit/update-visit/{id}`  
-**Méthode**: `PUT`  
- **Corps de la requête**:
  ```json
  {
    "visit_date": "2023-10-02",
    "visit_veto": "Nouveau vétérinaire",
    "visit_motif": "Nouveau motif",
    "visit_id_cat": 1
  }
- **Réponse**:
  ```json
  {
  "visit_date": "2023-10-02",
  "visit_veto": "Nouveau vétérinaire",
  "visit_motif": "Nouveau motif",
  "visit_id_cat": 1
  }

## Supprimer une visite
-**URL**: `/api/v1/clinique/visit/delete-visit/{id}`  
-**Méthode**: `DELETE`  
- **Réponse**:
  ```json
  {
    "message": "Vous avez supprimé la visite !"
  }

## Traitements

### Créer un traitement
-**URL**: `/api/v1/clinique/treatment/treatment_create`  
-**Méthode**: `POST`  
- **Corps de la requête**:
  ```json
  {
    "treatment_medoc": "Nom du médicament",
    "treatment_id_visit": 1
  }
- **Réponse**:
  ```json
  {
  "treatment_medoc": "Nom du médicament",
  "treatment_id_visit": 1
  }

## Obtenir tous les traitements
-**URL**: `/api/v1/clinique/treatment/all-treatments`  
-**Méthode**: `GET`  
- **Réponse**:
  ```json
  [
    {
      "treatment_medoc": "Nom du médicament",
      "treatment_id_visit": 1
    },
    ...
  ]

## Obtenir un traitement par ID
-**URL**: `/api/v1/clinique/treatment/one-treatment/{id}`  
-**Méthode**: `GET`  
- **Réponse**:
  ```json
  {
    "treatment_medoc": "Nom du médicament",
    "treatment_id_visit": 1
  }

## Mettre à jour un traitement
-**URL**: `/api/v1/clinique/treatment/update-treatment/{id}`  
-**Méthode**: `PUT`  
- **Corps de la requête**:
  ```json
  {
    "treatment_medoc": "Nouveau médicament",
    "treatment_id_visit": 1
  }
- **Réponse**:
  ```json
  {
  "treatment_medoc": "Nouveau médicament",
  "treatment_id_visit": 1
  }

## Supprimer un traitement
-**URL**: `/api/v1/clinique/treatment/delete-treatment/{id}`
-**Méthode**: `DELETE`
- **Réponse**:
  ```json
  {
  "message": "Vous avez supprimé un traitement !"
  }

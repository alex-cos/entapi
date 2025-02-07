// nolint: misspell
package entapi

type Entreprise struct {
	Siren                       string      `json:"siren"`
	NomComplet                  string      `json:"nom_complet"`
	RaisonSociale               string      `json:"nom_raison_sociale"`
	Sigle                       string      `json:"sigle"`
	NombreEtablissements        int32       `json:"nombre_etablissements"`
	NombreEtablissementsOuverts int32       `json:"nombre_etablissements_ouverts"`
	ActivitePrincipale          string      `json:"activite_principale"`
	CategorieEntreprise         string      `json:"categorie_entreprise,omitempty"`
	CaractereEmployeur          string      `json:"caractere_employeur,omitempty"`
	AnneCategorieEntreprise     string      `json:"annee_categorie_entreprise,omitempty"`
	DateCreation                string      `json:"date_creation"`
	DateMiseAJour               string      `json:"date_mise_a_jour,omitempty"`
	DateMiseAJourINSEE          string      `json:"date_mise_a_jour_insee,omitempty"`
	DateMiseAJourRNE            string      `json:"date_mise_a_jour_rne,omitempty"`
	EtatAdministratif           string      `json:"etat_administratif,omitempty"`
	NatureJuridique             string      `json:"nature_juridique,omitempty"`
	SectionActivitePrincipale   string      `json:"section_activite_principale,omitempty"`
	TrancheEffectifSalarie      string      `json:"tranche_effectif_salarie,omitempty"`
	AnneeTrancheEffectifSalarie string      `json:"annee_tranche_effectif_salarie,omitempty"`
	StatutDiffusion             string      `json:"statut_diffusion,omitempty"`
	Complements                 Complements `json:"complements,omitempty"`
	Siege                       Siege       `json:"siege,omitempty"`
	Dirigeants                  []Dirigeant `json:"dirigeants,omitempty"`
}

type Complements struct {
	CollectiviteTerritoriale       string `json:"collectivite_territoriale,omitempty"`
	ConventionCollectiveRenseignee bool   `json:"convention_collective_renseignee,omitempty"`
	EgaproRenseignee               bool   `json:"egapro_renseignee"`
	EstAssociation                 bool   `json:"est_association"`
	EstBio                         bool   `json:"est_bio"`
	EstEntrepreneurIndividuel      bool   `json:"est_entrepreneur_individuel"`
	EstEntrepreneurSpectacle       bool   `json:"est_entrepreneur_spectacle"`
	EstESS                         bool   `json:"est_ess"`
	EstFiness                      bool   `json:"est_finess"`
	EstOrganismeFormation          bool   `json:"est_organisme_formation"`
	EstQualiopi                    bool   `json:"est_qualiopi"`
	EstRGE                         bool   `json:"est_rge"`
	EstServicePublic               bool   `json:"est_service_public"`
	EstSiae                        bool   `json:"est_siae"`
	EstSocieteMission              bool   `json:"est_societe_mission"`
	EstUAI                         bool   `json:"est_uai"`
	IdentifiantAssociation         string `json:"identifiant_association,omitempty"`
}

type Siege struct {
	ActivitePrincipale               string `json:"activite_principale"`
	ActivitePrincipaleRegistreMetier string `json:"activite_principale_registre_metier,omitempty"`
	Adresse                          string `json:"adresse"`
	Cedex                            string `json:"cedex,omitempty"`
	CodePaysEtranger                 string `json:"code_pays_etranger,omitempty"`
	CodePostal                       string `json:"code_postal"`
	Commune                          string `json:"commune"`
	ComplementAdresse                string `json:"complement_adresse,omitempty"`
	Coordonnees                      string `json:"coordonnees"`
	DateCreation                     string `json:"date_creation"`
	DateDebutActivite                string `json:"date_debut_activite"`
	DateFermeture                    string `json:"date_fermeture,omitempty"`
	DateMiseAJour                    string `json:"date_mise_a_jour,omitempty"`
	DateMiseAJourINSEE               string `json:"date_mise_a_jour_insee,omitempty"`
	Departement                      string `json:"departement"`
	EPCI                             string `json:"epci,omitempty"`
	EstSiege                         bool   `json:"est_siege"`
	EtatAdministratif                string `json:"etat_administratif"`
	GeoAdresse                       string `json:"geo_adresse"`
	GeoID                            string `json:"geo_id"`
	Latitude                         string `json:"latitude"`
	Longitude                        string `json:"longitude"`
	NumeroVoie                       string `json:"numero_voie"`
	Region                           string `json:"region"`
	Siret                            string `json:"siret"`
	StatutDiffEtablissement          string `json:"statut_diffusion_etablissement"`
	TrancheEffectifSalarie           string `json:"tranche_effectif_salarie"`
	TypeVoie                         string `json:"type_voie"`
	LibelleCedex                     string `json:"libelle_cedex,omitempty"`
	LibelleCommune                   string `json:"libelle_commune"`
	LibelleCommuneEtranger           string `json:"libelle_commune_etranger,omitempty"`
	LibellePaysEtranger              string `json:"libelle_pays_etranger,omitempty"`
	LibelleVoie                      string `json:"libelle_voie,omitempty"`
}

type Dirigeant struct {
	Nom              string `json:"nom"`
	Prenoms          string `json:"prenoms"`
	AnneeDeNaissance string `json:"annee_de_naissance,omitempty"`
	DateDeNaissance  string `json:"date_de_naissance,omitempty"`
	Qualite          string `json:"qualite"`
	Nationalite      string `json:"nationalite"`
	TypeDirigeant    string `json:"type_dirigeant"`
}

type EntreprisesResponse struct {
	Results      []Entreprise `json:"results"`
	TotalResults int32        `json:"total_results"`
	Page         int32        `json:"page"`
	PerPage      int32        `json:"per_page"`
	TotalPages   int32        `json:"total_pages"`
}
